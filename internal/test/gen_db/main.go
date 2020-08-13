package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
)

func main() {
	f, err := os.Open("db.json")
	if err != nil {
		log.Fatalf("Can not open db.json: %v", err)
	}
	defer f.Close()

	var ad allData
	if err := json.NewDecoder(f).Decode(&ad); err != nil {
		log.Fatalf("Can not decode db.json: %v", err)
	}

	if err := writeFile(os.Stdout, ad.db); err != nil {
		log.Fatalf("Can not write output: %v", err)
	}
}

type allData struct {
	db map[string]json.RawMessage
}

func (db *allData) UnmarshalJSON(data []byte) error {
	var collections map[string][]json.RawMessage

	if err := json.Unmarshal(data, &collections); err != nil {
		return fmt.Errorf("decoding big data: %w", err)
	}

	db.db = make(map[string]json.RawMessage)
	for c, elements := range collections {
		for i, e := range elements {
			var withID struct {
				ID int `json:"id"`
			}
			if err := json.Unmarshal(e, &withID); err != nil {
				return fmt.Errorf("decoding id of collection %s, nr %d in list: %w", c, i, err)
			}
			db.db[fmt.Sprintf("%s:%d", c, withID.ID)] = e
		}
	}
	return nil
}

const tpl = `// Code generated with db.json DO NOT EDIT.
package test

import "encoding/json"

var exampleData = map[string]json.RawMessage{
	{{- range $key, $value := .Def}}
	"{{$key}}": []byte({{$.Escape}}{{$value}}{{$.Escape}}),
	{{- end}}
}
`

// The output needs the backtick (`) to work. But a backtick can not be used in
// a backtick-string. Therefore we use another byte and replace this byte with a
// backtick afterwards.
const escape byte = 1

func writeFile(w io.Writer, rlist map[string]json.RawMessage) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	strs := make(map[string]string, len(rlist))
	for k, v := range rlist {
		strs[k] = string(v)
	}

	data := map[string]interface{}{
		"Escape": string(escape),
		"Def":    strs,
	}

	if err := t.Execute(replacer{w: w}, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}

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
