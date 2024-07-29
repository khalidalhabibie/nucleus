package main

import (
	"testing"
)

func TestTimeConversion(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"evening time", "07:05:45PM", "19:05:45"},
		{"midnight", "12:01:00AM", "00:01:00"},
		{"noon", "12:01:00PM", "12:01:00"},
		{"morning", "01:00:00AM", "01:00:00"},
	}

	for _, testCase := range testCases {
		result := timeConversion(testCase.input)
		if result != testCase.expected {
			t.Errorf("%s failed: expected %s, got %s", testCase.name, testCase.expected, result)
		}
	}
}
