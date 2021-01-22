package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
	"github.com/ostcar/topic"
)

// Projectors holds the data of all projectors.
type Projectors struct {
	mu         sync.RWMutex
	projectors map[int]json.RawMessage
	closed     <-chan struct{}
	callables  map[string]projector.Callable
	topic      *topic.Topic
	ds         projector.Datastore
}

// NewProjectors returns a new projector instance.
func NewProjectors(ds projector.Datastore, ps map[string]projector.Callable, closed <-chan struct{}) *Projectors {
	return &Projectors{ds: ds, callables: ps, closed: closed}
}

// ProjectorData returns the data of every changed projector.
//
// Blocks untils the service or the context is closed or at least one projector
// has changed.
func (p *Projectors) ProjectorData(ctx context.Context, tid uint64) (uint64, map[int]json.RawMessage, error) {
	ntid, changedIDs, err := p.topic.Receive(ctx, tid)
	if err != nil {
		return 0, nil, fmt.Errorf("receiving from projector topic: %w", err)
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	data := make(map[int]json.RawMessage, len(changedIDs))
	for _, rid := range changedIDs {
		id := int(rid[0])
		data[id] = p.projectors[id]
	}
	return ntid, data, nil
}

// Update updates the cache of the projector.
func (p *Projectors) Update(data map[string]json.RawMessage) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.topic == nil {
		p.topic = topic.New(topic.WithClosed(p.closed))
		p.projectors = make(map[int]json.RawMessage)
	}

	for k, v := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		if parts[0] != "core/projector" {
			continue
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("key %s has wrong format. Expected id to be int not %s", k, parts[1])
		}

		if v == nil {
			// Projector was deleted.
			// TODO: How inform clients about a deleted projectors??
			delete(p.projectors, id)
			continue
		}

		if _, ok := p.projectors[id]; !ok {
			p.projectors[id] = nil
		}
	}

	var changed []string
	for id := range p.projectors {
		var elements struct {
			Elements []json.RawMessage `json:"elements"`
		}

		if err := p.ds.Get("core/projector", id, &elements); err != nil {
			return fmt.Errorf("decoding projector elements: %w", err)
		}

		ped := make([]projectorElementData, len(elements.Elements))
		for i, element := range elements.Elements {
			ped[i].Element = element
			var namer struct {
				Name string `json:"name"`
			}
			if err := json.Unmarshal(element, &namer); err != nil {
				if err := ped[i].setError(projector.NewClientError("unknwown slide None")); err != nil {
					return err
				}
				continue
			}

			c, ok := p.callables[namer.Name]
			if !ok {
				if namer.Name == "" {
					namer.Name = "None"
				}
				if err := ped[i].setError(projector.NewClientError("unknwown slide %s", namer.Name)); err != nil {
					return err
				}
				continue
			}

			data, err := c.Build(p.ds, element, id)
			if err != nil {
				if err := ped[i].setError(fmt.Errorf("building slide with %s: %w", namer.Name, err)); err != nil {
					return err
				}
				continue
			}

			ped[i].Data = data
		}

		rendered, err := json.Marshal(ped)
		if err != nil {
			return fmt.Errorf("decoding projector elements %w", err)
		}

		if bytes.Equal(p.projectors[id], rendered) {
			continue
		}

		p.projectors[id] = rendered
		changed = append(changed, string(rune(id)))
	}

	if len(changed) == 0 {
		// No changed data.
		return nil
	}

	log.Println("Projector-Data changed: ", changed)
	p.topic.Publish(changed...)
	return nil
}

type projectorData struct {
	elements []json.RawMessage
	rendered json.RawMessage
}

type projectorElementData struct {
	Element json.RawMessage `json:"element"`
	Data    json.RawMessage `json:"data"`
}

func (p *projectorElementData) setError(err error) error {
	var ce projector.ClientError
	if !errors.As(err, &ce) {
		p.Data = []byte(`{"error":"Internal error"}`)
		return err
	}

	data := struct {
		Error string `json:"error"`
	}{
		ce.Error(),
	}
	v, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encoding error message: %w", err)
	}
	p.Data = v
	return nil

}
