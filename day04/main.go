package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var lines []string
var matchingNumbersArray []int
var reNum = regexp.MustCompile(`\d+`)

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
	matchingNumbersArray = make([]int, len(lines))
}

func part1() {
	sumOfCardValues := 0
	start := time.Now()

	for cardIndex, card := range lines {
		cardValue := 0

		cardSplit := strings.Split(card[8:], " | ")
		numbersWeHave := reNum.FindAllString(cardSplit[1], -1)

		for _, num := range numbersWeHave {
			if isFound, _ := regexp.MatchString(`\b`+num+`\b`, cardSplit[0]); isFound {
				if cardValue == 0 {
					cardValue++
				} else {
					cardValue *= 2
				}
			}
		}

		if cardValue == 0 {
			matchingNumbersArray[cardIndex] = 0
		} else {
			matchingNumbersArray[cardIndex] = int(math.Log2(float64(cardValue))) + 1
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

	cardNbArray := make([]int, len(lines))

	for cardNbIndex := 0; cardNbIndex < len(cardNbArray); cardNbIndex++ {
		if cardNbArray[cardNbIndex] == 0 {
			cardNbArray[cardNbIndex] = 1
		}

		for repetition := 1; repetition <= cardNbArray[cardNbIndex]; repetition++ {
			for iOffset := 1; iOffset <= matchingNumbersArray[cardNbIndex]; iOffset++ {
				if cardNbArray[cardNbIndex+iOffset] == 0 {
					cardNbArray[cardNbIndex+iOffset] += 2
				} else {
					cardNbArray[cardNbIndex+iOffset] += 1
				}
			}
		}

		totalCards += cardNbArray[cardNbIndex]
	}

	fmt.Println("\npart2:", totalCards)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed.Seconds(), "s")
}

func main() {
	part1()
	part2()
}
