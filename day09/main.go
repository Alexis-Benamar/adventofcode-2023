package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var lines []string
var valuesList [][]int

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
	for _, line := range lines {
		numbersAsStr := strings.Split(line, " ")
		valuesToAdd := make([]int, len(numbersAsStr))
		for i, numAsStr := range numbersAsStr {
			valuesToAdd[i], _ = strconv.Atoi(string(numAsStr))
		}
		valuesList = append(valuesList, valuesToAdd)
	}
}

func generateDiffGrid(values []int) []int {
	diff := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		diff[i] = values[i+1] - values[i]
	}
	return diff
}

func isAllZeros(values []int) bool {
	allZeros := true
	for _, value := range values {
		if value != 0 {
			allZeros = false
			break
		}
	}

	return allZeros
}

func main() {
	start := time.Now()

	var sumOfNextValues, sumOfPreviousValues int

	for _, values := range valuesList {
		diffGrid := [][]int{values}
		currentDiff := generateDiffGrid(values)

		keepGoing := true
		for keepGoing {
			isLastDiff := isAllZeros(currentDiff)
			if isLastDiff {
				currentDiff = append(currentDiff, 0)
				keepGoing = false
			}

			diffGrid = append(diffGrid, currentDiff)

			if !isLastDiff {
				currentDiff = generateDiffGrid(currentDiff)
			}
		}

		// Part 2, extrapolating backwards
		for i := len(diffGrid)-1; i > 0; i-- {
			diffGrid[i - 1] = append([]int{diffGrid[i - 1][0] - diffGrid[i][0]}, diffGrid[i-1]...)

			// Add extrapolated value
			if i - 1 == 0 {
				sumOfPreviousValues += diffGrid[i - 1][0]
			}
		}

		// Sum of last digit of each row gives next number
		// Idea from this comment https://www.reddit.com/r/adventofcode/comments/18e5ytd/comment/kd8nbiu/
		nextValueToAdd := 0
		for _, values := range diffGrid {
			nextValueToAdd += values[len(values)-1]
		}

		sumOfNextValues += nextValueToAdd
	}

	fmt.Println("part1:", sumOfNextValues)
	fmt.Println("part2:", sumOfPreviousValues)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time %f s\n", elapsed.Seconds())
}
