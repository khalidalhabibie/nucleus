# Problem Solving Basic - Test 2 Solution

## Overview
This program calculates the ratios of positive, negative, and zero elements in an integer array and prints these ratios each formatted to six decimal places.

## Implementation
The function `plusMinus` takes an array of integers and computes the count of positive, negative, and zero values. It then calculates the ratio of each to the total number of elements in the array and prints these ratios with six decimal places.

## Running the Program
- Ensure you have Go installed on your machine.
- Place the source code in `main.go`.
- Run the program using:

```
go run main.go
```
## Testing
Unit tests are provided in `main_test.go`. To run these tests, ensure the test file is in the same directory as your main program and run:

```
go test
```


## Sample Output
Given the array `[-4, 3, -9, 0, 4, 1]`, the output will be:

```
0.500000
0.333333
0.166667
```

This output corresponds to the ratios of positive, negative, and zero values in the array, respectively.


