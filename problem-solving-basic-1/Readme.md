# MiniMaxSum Problem Solution

## Overview
This Golang program solves the "MiniMaxSum" problem where the task is to find the minimum and maximum sums from exactly four out of five integers.

## Sorting Algorithm
The program uses Quick Sort, a highly efficient sorting algorithm that operates in-place with a worst-case time complexity of O(n^2) and an average-case complexity of O(n log n). Quick Sort is a divide-and-conquer algorithm. It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot.

## How to Run
1. Ensure you have Go installed on your machine.
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program using the command:

``` bash
go run main.go
```



## Function Description
- `quickSort(arr []int, begin, end int)`: Sorts an array using the Quick Sort algorithm.
- `partition(arr []int, begin, end int) int`: Helps in partitioning the array for Quick Sort.
- `miniMaxSum(arr []int)`: Calculates and prints the minimum and maximum sums from the array.

The output will display the minimum sum and maximum sum calculated by adding four of the five integers.

## Example
For an input array `[9, 5, 7, 1, 3]`, the expected output after sorting and calculation would be:

```bash
16 24
```

This output represents the minimum and maximum sums respectively.

## Testing
It is recommended to test the program with different sets of integers to ensure its robustness. Here's a simple test case you can try:

Input: `[10, 2, 8, 4, 6]`
Expected Output: `20 28`

Adjust the test cases as needed to cover various scenarios including edge cases.

## Conclusion
This program is designed to provide a quick and efficient solution to the MiniMaxSum problem using the Quick Sort algorithm. It is suitable for educational purposes and real-world applications where similar problems are encountered.


