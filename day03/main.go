package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var lines []string
var gearsMap map[string][]int
var sumOfPartNumbers int
var sumOfGearRatios int
var reNum = regexp.MustCompile(`\d+`)
var reSymbol = regexp.MustCompile(`[^\d\s.]`)

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
	gearsMap = make(map[string][]int)
}

func checkPartNumber(partNumber, lineNumber, start, end int) (isPartNumber bool, symbol string, symX, symY int) {
	safeStart := start
	if start-1 <= 0 {
		safeStart = 0
	} else {
		safeStart = start - 1
	}

	safeEnd := end
	if end+1 >= len(lines[lineNumber]) {
		safeEnd = len(lines[lineNumber])
	} else {
		safeEnd = end + 1
	}

	for i := lineNumber - 1; i <= lineNumber+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}

		symRange := reSymbol.FindStringIndex(lines[i][safeStart:safeEnd])

		if symRange != nil {
			symbol := lines[i][safeStart+symRange[0] : safeStart+symRange[1]]

			return true, symbol, safeStart + symRange[0], i
		}
	}

	return false, "", 0, 0
}

func main() {
	start := time.Now()

	// Code here
	gearsMap := make(map[string][]int)

	for y, line := range lines {
		numRanges := reNum.FindAllStringIndex(line, -1)

		for _, numRange := range numRanges {
			numAsInt, _ := strconv.Atoi(lines[y][numRange[0]:numRange[1]])
			isPartNumber, symbol, symX, symY := checkPartNumber(numAsInt, y, numRange[0], numRange[1])

			if isPartNumber {
				sumOfPartNumbers += numAsInt

				if symbol == "*" {
					key := fmt.Sprintf("%d%d", symY, symX)
					if len(gearsMap[key]) > 0 {
						sumOfGearRatios += gearsMap[key][0] * numAsInt
					} else {
						gearsMap[key] = append(gearsMap[key], numAsInt)
					}
				}
			}
		}
	}

	fmt.Println("part1:", sumOfPartNumbers)
	fmt.Println("part2:", sumOfGearRatios)

	elapsed := time.Since(start)
	fmt.Println("\nExecution time", elapsed.Seconds(), "s")
}
