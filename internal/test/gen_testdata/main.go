package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

const (
	pathInput    = "export.json"
	pathFullData = "export.json.go"
	pathAppData  = "data_test.go"
)

func main() {
	f, err := os.Open(pathInput)
	if err != nil {
		log.Fatalf("Can not open db.json: %v", err)
	}
	defer f.Close()

	var e export
	e.All = make(map[string]json.RawMessage)
	e.Restricted = make(map[int]allData)
	if err := json.NewDecoder(f).Decode(&e); err != nil {
		log.Fatalf("Can not decode db.json: %v", err)
	}

	if err := writeData(e); err != nil {
		log.Fatalf("Can not write output: %v", err)
	}
}

type export struct {
	All           allData    `json:"all_data"`
	Restricted    restricted `json:"restricted_data"`
	RequiredUsers map[string]struct {
		IDs  []int  `json:"ids"`
		Perm string `json:"perm"`
	} `json:"required_users"`
	Projector []struct {
		Data    rawString `json:"data"`
		Element rawString `json:"element"`
	} `json:"projectors"`
}

type rawString struct {
	str string
}

func (s *rawString) UnmarshalJSON(data []byte) error {
	var b json.RawMessage
	if err := json.Unmarshal(data, &b); err != nil {
		return fmt.Errorf("decoding raw string: %w", err)
	}

	s.str = string(b)
	return nil
}

func (s *rawString) String() string {
	return string(s.str)
}

type allData map[string]json.RawMessage

func (ad allData) UnmarshalJSON(data []byte) error {
	var collections map[string][]json.RawMessage

	if err := json.Unmarshal(data, &collections); err != nil {
		return fmt.Errorf("decoding all data: %w", err)
	}

	for c, elements := range collections {
		for i, e := range elements {
			var withID struct {
				ID int `json:"id"`
			}
			if err := json.Unmarshal(e, &withID); err != nil {
				return fmt.Errorf("decoding id of collection %s, nr %d in list: %w", c, i, err)
			}
			ad[fmt.Sprintf("%s:%d", c, withID.ID)] = e
		}
	}
	return nil
}

type restricted map[int]allData

func (r restricted) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("decoding restricted data: %w", err)
	}

	for k, v := range raw {
		id, err := strconv.Atoi(k)
		if err != nil {
			return fmt.Errorf("Key is not an int %s: %w", k, err)
		}

		var ad allData = make(map[string]json.RawMessage)
		if err := json.Unmarshal(v, &ad); err != nil {
			return fmt.Errorf("decoding all data for user %d: %w", id, err)
		}
		r[id] = ad
	}

	return nil
}

const tplFullData = `// Code generated with export.json DO NOT EDIT.

package test

import "encoding/json"

var exampleData = map[string]json.RawMessage{
	{{- range $key, $value := .ExampleData}}
	"{{$key}}": []byte({{$.Escape}}{{$value}}{{$.Escape}}),
	{{- end}}
}

var exampleRestrictedData = map[int]map[string]json.RawMessage{
	{{- range $uid, $data := .ExampleRestrictedData}}
	{{$uid}}: {
		{{- range $key, $value := $data}}
		"{{$key}}": []byte({{$.Escape}}{{$value}}{{$.Escape}}),
		{{- end}}
	},
	{{- end}}
}

type permIDs struct{
	perm string
	ids  []int
}

var exampleRequiredUser = map[string]permIDs{
	{{- range $key, $data := .ExampleRequiredUser}}
	"{{$key}}": {
		"{{$data.Perm}}",
		[]int{ {{$data.IDs}} },
	},
	{{- end}}
}

type projectorData struct {
	Data    json.RawMessage
	Element json.RawMessage
}

var exampleProjector = []projectorData{
	{{- range $row := .ExampleProjector }}
	{
		[]byte({{$.Escape}}{{$row.Data}}{{$.Escape}}),
		[]byte({{$.Escape}}{{$row.Element}}{{$.Escape}}),
	},
	{{- end}}
}
`

func writeData(e export) error {
	f, err := os.Create(pathFullData)
	if err != nil {
		log.Fatalf("Can not create %s: %v", pathFullData, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Can not close %s: %v", pathFullData, err)
		}
	}()

	t, err := template.New("t").Parse(tplFullData)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	strExampleData := make(map[string]string, len(e.All))
	for k, v := range e.All {
		strExampleData[k] = string(v)
	}

	strRestrictedData := make(map[int]map[string]string, len(e.Restricted))
	for uid, data := range e.Restricted {
		strRestrictedData[uid] = make(map[string]string)
		for k, v := range data {
			strRestrictedData[uid][k] = string(v)
		}
	}

	strRequiredUser := make(map[string]struct {
		Perm string
		IDs  string
	}, len(e.RequiredUsers))
	for k, data := range e.RequiredUsers {
		strRequiredUser[k] = struct {
			Perm string
			IDs  string
		}{
			data.Perm,
			intsToStr(data.IDs),
		}
	}

	data := map[string]interface{}{
		"Escape":                string(escape),
		"ExampleData":           strExampleData,
		"ExampleRestrictedData": strRestrictedData,
		"ExampleRequiredUser":   strRequiredUser,
		"ExampleProjector":      e.Projector,
	}

	if err := t.Execute(replacer{w: f}, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}

// The output needs the backtick (`) to work. But a backtick can not be used in
// a backtick-string. Therefore we use another byte and replace this byte with a
// backtick afterwards.
const escape byte = 1

type replacer struct {
	w io.Writer
}

func (r replacer) Write(p []byte) (n int, err error) {
	for i, b := range p {
		if b == escape {
			p[i] = '`'
		}
	}
	return r.w.Write(p)
}

func intsToStr(ints []int) string {
	var s string
	for _, i := range ints {
		s += fmt.Sprintf("%d,", i)
	}
	return strings.TrimSuffix(s, ",")
}
