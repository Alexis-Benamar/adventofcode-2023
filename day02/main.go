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

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}
}

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var sumOfIndeces int
var sumOfPowers int

func main() {
	start := time.Now()

	// Code here
	lines := strings.Split(data, "\n")

	reRed := regexp.MustCompile(`(\d+)?(\sred)+`)
	reGreen := regexp.MustCompile(`(\d+)?(\sgreen)+`)
	reBlue := regexp.MustCompile(`(\d+)?(\sblue)+`)

	for lineIndex, line := range lines {
		sets := strings.Split(line[8:], "; ")
		shouldAdd := true
		var minRed int
		var minGreen int
		var minBlue int

		for _, set := range sets {
			var red int
			var green int
			var blue int

			redMatch := reRed.FindStringSubmatch(set)
			if redMatch != nil {
				red, _ = strconv.Atoi(redMatch[1])

				if red > minRed {
					minRed = red
				}
			}

			greenMatch := reGreen.FindStringSubmatch(set)
			if greenMatch != nil {
				green, _ = strconv.Atoi(greenMatch[1])

				if green > minGreen {
					minGreen = green
				}
			}

			blueMatch := reBlue.FindStringSubmatch(set)
			if blueMatch != nil {
				blue, _ = strconv.Atoi(blueMatch[1])

				if blue > minBlue {
					minBlue = blue
				}
			}

			if red > MAX_RED || green > MAX_GREEN || blue > MAX_BLUE {
				shouldAdd = false
			}
		}

		if shouldAdd {
			sumOfIndeces += lineIndex + 1
		}

		powerOfSet := minRed * minGreen * minBlue

		sumOfPowers += powerOfSet
	}

	fmt.Println("part1:", sumOfIndeces)
	fmt.Println("part2:", sumOfPowers)

	elapsed := time.Since(start)
	fmt.Println("\nExecution time", elapsed.Seconds(), "s")
}
