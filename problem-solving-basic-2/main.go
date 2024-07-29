package main

import (
	"fmt"
)

// plusMinus calculates and prints the ratios of positive, negative, and zero values in the array.
func plusMinus(arr []int) {
	n := len(arr)
	positiveCount, negativeCount, zeroCount := 0, 0, 0

	for _, value := range arr {
		switch {
		case value > 0:
			positiveCount++
		case value < 0:
			negativeCount++
		case value == 0:
			zeroCount++
		}
	}

	fmt.Printf("%.6f\n", float64(positiveCount)/float64(n))
	fmt.Printf("%.6f\n", float64(negativeCount)/float64(n))
	fmt.Printf("%.6f\n", float64(zeroCount)/float64(n))
}

func main() {
	arr := []int{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
