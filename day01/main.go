package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Change input.txt content between part 1 & part 2
//
//go:embed input.txt
var data string

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}
}

// Reverse a string
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func parseInput(outputString string, lines []string, firstRegex, lastRegex *regexp.Regexp) {
	var res int

	numAsStrings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		reversedLine := reverse(line)

		first := firstRegex.FindString(line)
		if _, err := strconv.Atoi(first); err != nil {
			first = numAsStrings[first]
		}

		last := lastRegex.FindString(reversedLine)
		if _, err := strconv.Atoi(last); err != nil {
			last = numAsStrings[reverse(last)]
		}

		numToAdd, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

		// Debugging
		// fmt.Println(line, first, last, numToAdd)

		res += numToAdd
	}

	fmt.Printf("\n%s: %d", outputString, res)
}

func main() {
	start := time.Now()

	// Code here
	lines := strings.Split(data, "\n")

	// Part 1

	reNumOnly := regexp.MustCompile(`\d`)
	parseInput("part1", lines, reNumOnly, reNumOnly)

	// Part 2

	plainTextNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	plainTextNumbersReverse := make([]string, len(plainTextNumbers))
	for i, num := range plainTextNumbers {
		plainTextNumbersReverse[i] = reverse(num)
	}

	re := regexp.MustCompile(`(\d|` + strings.Join(plainTextNumbers, "|") + `)`)
	reReverse := regexp.MustCompile(`(\d|` + strings.Join(plainTextNumbersReverse, "|") + `)`)

	parseInput("part2", lines, re, reReverse)

	elapsed := time.Since(start)
	fmt.Println("\nExecution time", elapsed.Seconds(), "s")
}
