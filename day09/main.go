package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed example.txt
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
	diff := make([]int, len(values) - 1)
	for i := 0; i < len(values) - 1; i++ {
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

	sumOfNextValues := 0

	for _, values := range valuesList {
		diffGrid := [][]int{values}
		currentDiff := generateDiffGrid(values)

		keepGoing := true
		for keepGoing {
			isLastDiff := isAllZeros(currentDiff)
			if isLastDiff {
				keepGoing = false
			}

			diffGrid = append(diffGrid, currentDiff)

			if !isLastDiff {
				currentDiff = generateDiffGrid(currentDiff)
			}
		}

		fmt.Println("\n", diffGrid)

		for i := len(diffGrid) - 1; i >= 0; i-- {
			if diffGrid[i][len(diffGrid[i]) - 1] == 0 {
				continue
			}

			lastOfPrevRow := diffGrid[i+1][len(diffGrid[i+1]) - 1]
			lastOfCurrent := diffGrid[i][len(diffGrid[i]) - 1]

			fmt.Println("adding", lastOfCurrent + lastOfPrevRow)

			diffGrid[i] = append(diffGrid[i], lastOfCurrent + lastOfPrevRow)
		}

		newValue := diffGrid[0][len(diffGrid[0]) - 1]
		sumOfNextValues += newValue
		fmt.Println("new value", newValue)
	}

	fmt.Println("\npart1:", sumOfNextValues)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}
