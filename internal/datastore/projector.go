package datastore

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ostcar/topic"
)

type projector struct {
	mu         sync.RWMutex
	projectors map[int]*projectorData
	closed     <-chan struct{}
	callables  map[string]ProjectorCallable
}

func (p *projector) update(data map[string]json.RawMessage) error {
	p.mu.Lock()
	defer p.mu.Unlock()

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
			// New projector, create it.
			p.projectors[id] = &projectorData{
				topic: topic.New(topic.WithClosed(p.closed)),
			}
		}

		// Update projector
		var elements struct {
			Elements []json.RawMessage `json:"elements"`
		}
		if err := json.Unmarshal(v, &elements); err != nil {
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
			}

			data, err := c.Build(element, id)
			if err != nil {
				ped[i].Error = err.Error()
				continue
			}

			ped[i].Element = element
			ped[i].Data = data
		}

		pData, err := json.Marshal(ped)
		if err != nil {
			return fmt.Errorf("decoding projector elements %w", err)
		}

		p.projectors[id].elements = pData
		p.projectors[id].topic.Publish()
	}
	return nil
}

type projectorData struct {
	topic    *topic.Topic
	elements json.RawMessage
}

type projectorElementData struct {
	Element json.RawMessage `json:"element"`
	Data    json.RawMessage `json:"data"`
	Error   string          `json:"error,omitempty"`
}
