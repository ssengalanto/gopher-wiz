package main

import (
	"reflect"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  []int
	}{
		{name: "match found", input: Input{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "no match found", input: Input{nums: []int{2, 7, 11, 15}, target: 6}, want: nil},
		{name: "empty slice", input: Input{nums: []int{}, target: 26}, want: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := twoSum(test.input.nums, test.input.target)

			if isMatched := reflect.DeepEqual(got, test.want); !isMatched {
				t.Errorf("\ntwoSum(%v, %d) \ngot = %v \nwant = %v", test.input.nums, test.input.target, got, test.want)
			}
		})
	}
}
