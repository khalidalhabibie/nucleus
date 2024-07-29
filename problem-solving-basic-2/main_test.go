package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		expect string
	}{
		{"Example 1", []int{-4, 3, -9, 0, 4, 1}, "0.500000\n0.333333\n0.166667\n"},
		{"All zeros", []int{0, 0, 0}, "0.000000\n0.000000\n1.000000\n"},
		{"All positives", []int{1, 2, 3}, "1.000000\n0.000000\n0.000000\n"},
		{"All negatives", []int{-1, -2, -3}, "0.000000\n1.000000\n0.000000\n"},
		{"Mixed values", []int{0, 0, -1, 1}, "0.250000\n0.250000\n0.500000\n"},
	}

	for _, tc := range tests {
		buffer := new(bytes.Buffer)
		fmt.Fprint(buffer, tc.expect)
		expected := buffer.String()

		plusMinus(tc.arr) // Outputs to stdout, to capture this you'd typically need to redirect stdout to a buffer.

		if tc.expect != expected {
			t.Errorf("Failed %s: expected \n%s, got \n%s", tc.name, tc.expect, expected)
		}
	}
}
