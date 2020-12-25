package test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func Benchmark_encode_string_with_SetEscapeHTML(b *testing.B) {
	type V struct {
		S string
		B bool
		I int
		O struct {
			a string
			c bool
			i int
		}
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// buf := &bytes.Buffer{}
		// enc := json.NewEncoder(buf)
		// enc.SetEscapeHTML(true)
		if buf, err := json.Marshal(V{S: "s", B: true, I: 233, O: struct {
			a string
			c bool
			i int
		}{a: "dfsafdsa", c: true, i: 1233466}}); err != nil {
			b.Fatal(buf)
		}
	}
}

func Benchmark_encode_string(b *testing.B) {
	type V struct {
		S string
		B bool
		I int
		O struct {
			a string
			c bool
			i int
		}
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if buf, err := json.Marshal(V{S: "s", B: true, I: 233, O: struct {
			a string
			c bool
			i int
		}{a: "dfsafdsa", c: true, i: 1233466}}); err != nil {
			b.Fatal(buf)
		}
	}
}
