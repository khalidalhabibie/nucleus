# Time Conversion Solution

## Overview
This program converts a time given in 12-hour AM/PM format to 24-hour military format.

## Functionality
The function `timeConversion` takes a string representing a time in 12-hour format and converts it to a string representing the time in 24-hour format.

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
Given the time `07:05:45PM`, the output will be:

```
19:05:45
```

Given the time `12:01:00AM`, the output will be:

```
00:01:00
```


## Conclusion
This program uses simple string manipulation and condition checks to perform the conversion, making it efficient and straightforward to understand and use.
