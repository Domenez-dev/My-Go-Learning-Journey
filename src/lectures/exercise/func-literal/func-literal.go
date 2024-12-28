//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import "fmt"

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	line_func := func(line string, count map[string]int) {
		for _, r := range line {
			switch {
			case r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z':
				count["letters"]++
			case r >= '0' && r <= '9':
				count["digits"]++
			case r == ' ':
				count["spaces"]++
			case r == '.' || r == ',' || r == '!' || r == '?':
				count["puncts"]++
			}
		}
	}

	count := map[string]int{
		"letters": 0,
		"digits":  0,
		"spaces":  0,
		"puncts":  0,
	}

	for _, line := range lines {
		line_func(line, count)
	}

	fmt.Printf("Letters: %d\n", count["letters"])
	fmt.Printf("Digits: %d\n", count["digits"])
	fmt.Printf("Spaces: %d\n", count["spaces"])
	fmt.Printf("Punctuation: %d\n", count["puncts"])
}
