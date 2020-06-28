package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
	"github.com/ostcar/topic"
)

type projectors struct {
	mu         sync.RWMutex
	projectors map[int]json.RawMessage
	closed     <-chan struct{}
	callables  map[string]projector.Callable
	topic      *topic.Topic
	ds         *Datastore
}

// Projectors returns the data of every changed projector.
//
// Blocks untils the service or the context is closed or at least one projector
// has changed.
func (p *projectors) Projectors(ctx context.Context, tid uint64) (uint64, map[int]json.RawMessage, error) {
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

func (p *projectors) update(data map[string]json.RawMessage) error {
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
			var namer struct {
				Name string `json:"name"`
			}
			if err := json.Unmarshal(element, &namer); err != nil {
				ped[i].Error = fmt.Sprintf("projector element has no name: %v", err)
				continue
			}

			c, ok := p.callables[namer.Name]
			if !ok {
				ped[i].Error = fmt.Sprintf("unknown projector element %s", namer.Name)
				continue
			}

			data, err := c.Build(p.ds, element, id)
			if err != nil {
				ped[i].Error = err.Error()
				continue
			}

			ped[i].Element = element
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
		changed = append(changed, string(id))
	}

	if len(changed) == 0 {
		// No changed data.
		return nil
	}

	p.topic.Publish(changed...)
	return nil
}

type projectorData struct {
	elements []json.RawMessage
	rendered json.RawMessage
}

type projectorElementData struct {
	Element json.RawMessage `json:"element,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}
