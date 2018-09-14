// +build ignore

//go:generate go run gen.go

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/benjaminestes/crawl/src/crawler/data"
)

type SchemaItem struct {
	Description string
	Mode        string
	Name        string
	Type        string
	Fields      []SchemaItem
}

func genFile(name string, buf *bytes.Buffer) {
	fmt.Println(buf)
	b, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := ioutil.WriteFile(name, b, 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "// Code generated by go generate gen.go; DO NOT EDIT.\n")
	fmt.Fprintln(&buf, "//go:generate go run gen.go\n")
	fmt.Fprintln(&buf, "package schema")

	r := &data.Result{}
	t := reflect.TypeOf(*r)

	fmt.Fprintln(&buf, "var BQ = []SchemaItem{")
	recursiveGenerate(t, &buf)
	fmt.Fprintln(&buf, "}")
	genFile("schema.go", &buf)
}

func recursiveGenerate(t reflect.Type, buf *bytes.Buffer) {
	for i := 0; i < t.NumField(); i++ {
		g := t.Field(i)
		s := &SchemaItem{}
		s.Name = g.Name
		s.Type = typeToString(g.Type)

		if isRepeated(g.Type) {
			s.Mode = "REPEATED"
		} else {
			s.Mode = defaultMode(g.Tag.Get("mode"))
		}

		fmt.Fprintf(buf, "{\nName:\"%s\",\nType:\"%s\",\nMode:\"%s\",\n", s.Name, s.Type, s.Mode)

		switch g.Type.Kind() {
		case reflect.Ptr:
			fmt.Fprintln(buf, "Fields: []SchemaItem{")
			recursiveGenerate(g.Type.Elem(), buf)
			fmt.Fprintln(buf, "},")
		case reflect.Slice:
			fmt.Fprintln(buf, "Fields: []SchemaItem{")
			recursiveGenerate(g.Type.Elem().Elem(), buf)
			fmt.Fprintln(buf, "},")
		}

		fmt.Fprintf(buf, "},\n")
	}
}

func defaultMode(s string) string {
	if s == "" {
		return "NULLABLE"
	}
	return s
}

func isRepeated(t reflect.Type) bool {
	if t.Kind() == reflect.Slice {
		return true
	}
	return false
}

func typeToString(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Float64:
		return "FLOAT64"
	case reflect.Int, reflect.Int64:
		return "INT64"
	case reflect.String:
		return "STRING"
	case reflect.Bool:
		return "BOOL"
	case reflect.Ptr:
		return "RECORD"
	case reflect.Slice:
		return typeToString(t.Elem())
	default:
		return "STRING"
	}
}
