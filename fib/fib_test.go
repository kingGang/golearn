package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d)=%d; expected=%d", in, actual, expected)
	}
}

func ExampleFib() {
	fmt.Println("OK")
}

func TestFib2(t *testing.T) {
	var fibTests = []struct {
		in       int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}
	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d)=%d; expected=%d", tt.in, actual, tt.expected)
		}
	}
}
