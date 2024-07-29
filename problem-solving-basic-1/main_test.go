package main

import (
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		wantMin int
		wantMax int
	}{
		{"sorted positive", []int{1, 2, 3, 4, 5}, 10, 14},
		{"unsorted positive", []int{5, 3, 1, 4, 2}, 10, 14},
		{"mixed numbers", []int{7, 69, 2, 221, 8974}, 299, 9271},
		// {"all negatives", []int{-10, -1, -3, -4, -5}, -18, -10},
	}

	for _, tt := range tests {
		// Copy the array to avoid modifying the original in tests
		arrCopy := make([]int, len(tt.arr))
		copy(arrCopy, tt.arr)

		miniMaxSum(arrCopy) // This prints the output and we assume it sorts the array correctly

		// Recalculate to check
		gotMin := arrCopy[0] + arrCopy[1] + arrCopy[2] + arrCopy[3]
		gotMax := arrCopy[1] + arrCopy[2] + arrCopy[3] + arrCopy[4]

		if gotMin != tt.wantMin || gotMax != tt.wantMax {
			t.Errorf("%s failed: expected (%d, %d), got (%d, %d)", tt.name, tt.wantMin, tt.wantMax, gotMin, gotMax)
		}
	}
}

func BenchmarkMiniMaxSum(b *testing.B) {
	arr := []int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		miniMaxSum(arrCopy)
	}
}
