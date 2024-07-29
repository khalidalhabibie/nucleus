package main

import (
	"fmt"
)

// Function to perform Quick Sort
func quickSort(arr []int, begin, end int) {
	if begin < end {
		pi := partition(arr, begin, end)
		quickSort(arr, begin, pi-1)
		quickSort(arr, pi+1, end)
	}
}

// Helper function for Quick Sort to find the partition position
func partition(arr []int, begin, end int) int {
	pivot := arr[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}

// Function to calculate miniMaxSum using the sorted array
func miniMaxSum(arr []int) {
	quickSort(arr, 0, len(arr)-1) // Sort the array using Quick Sort
	minSum := arr[0] + arr[1] + arr[2] + arr[3]
	maxSum := arr[1] + arr[2] + arr[3] + arr[4]
	fmt.Printf("%d %d\n", minSum, maxSum)
}

func main() {
	arr := []int{9, 5, 7, 1, 3}
	miniMaxSum(arr)
}
