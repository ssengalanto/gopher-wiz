package main

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "positive", input: 12345, want: 54321},
		{name: "negative", input: -12345, want: -54321},
		{name: "zero", input: 0, want: 0},
		{name: "exceeds max int32", input: math.MaxInt32 + 1, want: 0},
		{name: "leading zero/es", input: 200, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := reverse(test.input); got != test.want {
				t.Errorf("reverse(%d)\n got = %d\b want = %d\n", test.input, got, test.want)
			}
		})

	}
}
