package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var lines []string
var reNum = regexp.MustCompile(`\d+`)

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
}

func part1() {
	sumOfCardValues := 0
	start := time.Now()

	for _, line := range lines {
		cardValue := 0

		cardSplit := strings.Split(line[8:], " | ")
		numbersWeHave := reNum.FindAllString(cardSplit[1], -1)

		for _, num := range numbersWeHave {
			if isFound, _ := regexp.MatchString(`\b` + num + `\b`, cardSplit[0]); isFound {
				if cardValue == 0 {
					cardValue++
				} else {
					cardValue *= 2
				}
			}
		}

		sumOfCardValues += cardValue
	}

	fmt.Println("part1:", sumOfCardValues)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed.Seconds(), "s")
}

func part2() {
	start := time.Now()
	totalCards := 0

	fmt.Println("\npart2:", totalCards)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed.Seconds(), "s")
}

func main() {
	part1()
	part2()
}