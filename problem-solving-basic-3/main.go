package main

import (
	"fmt"
	"strconv"
	"strings"
)

// timeConversion converts a time from 12-hour AM/PM format to 24-hour military time.
func timeConversion(s string) string {
	pm := strings.Contains(s, "PM")
	parts := strings.Split(strings.TrimSuffix(strings.TrimSuffix(s, "AM"), "PM"), ":")
	hour, _ := strconv.Atoi(parts[0])
	minute := parts[1]
	second := parts[2]

	if pm {
		if hour != 12 {
			hour += 12
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}

	return fmt.Sprintf("%02d:%s:%s", hour, minute, second)
}

func main() {
	fmt.Println(timeConversion("07:05:45PM")) // Should print "19:05:45"
	fmt.Println(timeConversion("12:01:00AM")) // Should print "00:01:00"
}
