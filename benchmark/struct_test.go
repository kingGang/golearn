package mystruct

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BenchmarkToJson(b *testing.B) {
	now := time.Now()
	agent := &Agent{
		Id:            primitive.NewObjectID(),
		Name:          "wxg",
		UpLevelLessAt: &now,
	}
	for i := 0; i < b.N; i++ {
		toJson(agent)
	}
}

func BenchmarkToJson1(b *testing.B) {
	now := time.Now()
	agent := &struct {
		Phone         string
		Name          string
		UpLevelLessAt time.Time
		Address       string
		Age           int
		Dsitrb        string
		A             string
		B             string
		C             string
		D             string
		E             string
		F             string
		G             string
		H             int64
		I             struct{}
	}{
		Phone:         "16987655678",
		Name:          "wxg",
		UpLevelLessAt: now,
	}
	for i := 0; i < b.N; i++ {
		toJson(agent)
	}
}
