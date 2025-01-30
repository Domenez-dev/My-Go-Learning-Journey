//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	time string
	err  string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v, %v", t.time, t.err)
}

func TimeParser(input string) (Time, error) {
	time := strings.Split(input, ":")
	if len(time) != 3 {
		return Time{0, 0, 0}, &TimeParseError{input, "bad format type hh:mm:ss"}
	}

	// check hours
	hours, err := strconv.Atoi(time[0])
	if err != nil {
		return Time{0, 0, 0}, &TimeParseError{input, "error parsing integer type"}
	}

	if hours > 23 || hours < 00 {
		return Time{0, 0, 0}, &TimeParseError{input, "Hours must be < 24 and > 00"}
	}

	// check minutes
	minutes, err := strconv.Atoi(time[1])
	if err != nil {
		return Time{0, 0, 0}, &TimeParseError{input, "error parsing integer type"}
	}

	if minutes > 59 || minutes < 00 {
		return Time{0, 0, 0}, &TimeParseError{input, "minutes must be < 24 and > 00"}
	}

	// check seconds
	seconds, err := strconv.Atoi(time[2])
	if seconds > 59 || seconds < 00 {
		return Time{0, 0, 0}, &TimeParseError{input, "seconds must be < 24 and > 00"}
	}

	if err != nil {
		return Time{0, 0, 0}, &TimeParseError{input, "error parsing integer type"}
	}
	return Time{hours, minutes, seconds}, nil
}
