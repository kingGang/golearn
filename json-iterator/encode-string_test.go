package test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Benchmark_encode_string_with_SetEscapeHTML(b *testing.B) {
	type V struct {
		S string
		B bool
		I int
		O primitive.ObjectID
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// buf := &bytes.Buffer{}
		// enc := json.NewEncoder(buf)
		// enc.SetEscapeHTML(true)
		if buf, err := json.Marshal(V{S: "s", B: true, I: 233, O: primitive.NewObjectID()}); err != nil {
			b.Fatal(buf)
		}
	}
}

func Benchmark_encode_string(b *testing.B) {
	type V struct {
		S string
		B bool
		I int
		O string
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if buf, err := json.Marshal(V{S: "s", B: true, I: 233, O: "123456abc78931646543213246"}); err != nil {
			b.Fatal(buf)
		}
	}
}
