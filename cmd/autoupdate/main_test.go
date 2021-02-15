package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/datastore"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestRestrictedData(t *testing.T) {
	for _, tt := range test.ExampleRestrictedData() {
		t.Run(tt.Name, func(t *testing.T) {
			restricters := openslidesRestricters(tt.Permer)
			r := restricters[tt.Collection]

			got, err := r.Restrict(tt.UID, tt.Element)
			if err != nil {
				t.Errorf("Restrict returned unexpected error: %v ", err)
			}

			// Full restriction.
			if tt.Expected == nil {
				if got != nil {
					t.Errorf("Restrict returned `%s`, expected nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("Restrict() returned nil, expected %s", tt.Expected)
				return
			}

			test.ExpectEqualJSON(t, got, tt.Expected)
		})
	}
}

func TestRequiredUser(t *testing.T) {
	required := openslidesRequiredUsers()

	for _, tt := range test.ExampleRequiredUser() {
		t.Run(tt.Name, func(t *testing.T) {
			r, ok := required[tt.Collection]
			if !ok {
				// No required users for element.
				return
			}

			ids, perm, err := r(tt.Element)
			if err != nil {
				t.Errorf("RequiredUser returned unexpected error: %v", err)
			}

			if perm != tt.ExpectPerm {
				t.Errorf("RequiredUser returned perm %s, expected %s", perm, tt.ExpectPerm)
			}

			if !test.CmpIntSet(ids, tt.ExpectIDs) {
				t.Errorf("RequiredUser returned ids %v, expected %v", ids, tt.ExpectIDs)
			}
		})
	}
}

func TestProjector(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	callables := openslidesProjectorCallables()
	ds := test.NewDatastoreMock(0, closed)

	for _, tt := range test.ExampleProjector() {
		t.Run(tt.Name, func(t *testing.T) {
			fd := make(map[string]json.RawMessage)
			for k, v := range test.ExampleData() {
				fd[k] = v
			}
			for k, v := range tt.Overwrite {
				fd[k] = v
			}
			ds.FullData = fd

			p := datastore.NewProjectors(ds, callables, closed)
			if err := p.Update(fd); err != nil {
				t.Fatalf("Can not update projector data: %v", err)
			}

			_, projectors, err := p.ProjectorData(context.Background(), 0)
			if err != nil {
				t.Fatalf("Can not get projectors: %v", err)
			}

			var v []json.RawMessage
			if err := json.Unmarshal(projectors[1], &v); err != nil {
				t.Fatalf("Can not decode first projector: %v", err)
			}

			test.ExpectEqualJSON(t, v[0], tt.Expected)
		})
	}
}

func TestSecret(t *testing.T) {
	for _, tt := range []struct {
		name    string
		content string
		result  string
	}{
		{
			"single quotes",
			`DJANGO_SECRET_KEY='development'`,
			"development",
		},
		{
			"douple quotes",
			`DJANGO_SECRET_KEY="development'`,
			"development",
		},
		{
			"white spaces",
			`DJANGO_SECRET_KEY =	"development'`,
			"development",
		},
		{
			"multi line",
			`VALUE="foo"
			DJANGO_SECRET_KEY="development'
			OTHER="bar"`,
			"development",
		},
		{
			"quotes inside",
			`DJANGO_SECRET_KEY="devel"op'ment'`,
			"devel\"op'ment",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			secret, err := secretKey(strings.NewReader(tt.content))
			if tt.result != "" && err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			if secret != tt.result {
				t.Errorf("secredKey(%s) == `%s`; expected `%s`", tt.content, secret, tt.result)
			}
		})
	}
}
