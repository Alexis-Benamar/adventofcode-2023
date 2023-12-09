package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type HandType int

type Hand struct {
	Cards string
	Bid   int
	Type  HandType
}

const (
	HighCard    HandType = 0
	OnePair              = 1
	TwoPair              = 2
	ThreeOfKind          = 3
	FullHouse            = 4
	FourOfKind           = 5
	FiveOfKind           = 6
)

//go:embed input.txt
var data string
var lines []string
var hands, handsWithJoker []Hand

/*
Returns HandType by creating an array of different characters within cards
*/
func getHandType(cards string, withJoker bool) (handType HandType) {
	diffChars := make(map[string]int, len(cards))

	for _, char := range cards {
		if _, hasChar := diffChars[string(char)]; hasChar {
			diffChars[string(char)] += 1
		} else {
			diffChars[string(char)] = 1
		}
	}

	switch len(diffChars) {
	// Must be High Hand
	case 5:
		_, hasJ := diffChars["J"]

		if withJoker && hasJ {
			return HandType(1)
		} else {
			return HandType(0)
		}
	// Must be One Pair
	case 4:
		_, hasJ := diffChars["J"]

		if withJoker && hasJ {
			return HandType(3)
		} else {
			return HandType(1)
		}
	// If at least one value == 2, must be Two Pair
	// If at least one value == 3, must be Three of Kind
	case 3:
		_, hasJ := diffChars["J"]

		if withJoker && hasJ {
			if diffChars["J"] == 1 {
				values := []int{}
				for _, value := range diffChars {
					values = append(values, value)
				}

				if values[0] == 3 || values[1] == 3 || values[2] == 3 {
					return HandType(5)
				} else {
					return HandType(4)
				}
			} else {
				return HandType(5)
			}
		} else {
			values := []int{}
			for _, value := range diffChars {
				values = append(values, value)
			}

			if values[0] == 2 || values[1] == 2 || values[2] == 2 {
				return HandType(2)
			}

			if values[0] == 3 || values[1] == 3 || values[2] == 3 {
				return HandType(3)
			}
		}

	// If at least one value == 2, must be Full House
	// Else, must be Four of Kind
	case 2:
		values := []int{}
		for _, value := range diffChars {
			values = append(values, value)
		}

		_, hasJ := diffChars["J"]
		if withJoker && hasJ {
			return HandType(6)
		} else {
			if values[0] == 2 || values[1] == 2 {
				return HandType(4)
			} else {
				return HandType(5)
			}
		}
	// Must be Five of Kind
	default:
		return HandType(6)
	}

	return 0
}

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
	hands = make([]Hand, len(lines))
	handsWithJoker = make([]Hand, len(lines))

	for i, line := range lines {
		tokens := strings.Split(line, " ")
		bid, _ := strconv.Atoi(tokens[1])

		hand := Hand{
			Cards: tokens[0],
			Bid:   bid,
			Type:  getHandType(tokens[0], false),
		}

		handWithJoker := Hand{
			Cards: tokens[0],
			Bid:   bid,
			Type:  getHandType(tokens[0], true),
		}

		hands[i] = hand
		handsWithJoker[i] = handWithJoker
	}
}

func sortHands(handsToSort []Hand, cardsOrder string) []Hand {
	isDone := false
	sortedHands := make([]Hand, len(handsToSort))

	copy(sortedHands, handsToSort)

	// Bubble sort
	for !isDone {
		isDone = true
		handIndex := 0

		for handIndex < len(sortedHands)-1 {
			// If left hand's type is higher than right hand's type, swap
			if sortedHands[handIndex].Type > sortedHands[handIndex+1].Type {
				sortedHands[handIndex], sortedHands[handIndex+1] = sortedHands[handIndex+1], sortedHands[handIndex]
				isDone = false
			} else if sortedHands[handIndex].Type == sortedHands[handIndex+1].Type {
				// If types are the same, we loop over both hands's cards
				// If leftCard < rightCard, keep hand order
				// If leftCard > rightCard, swap
				// Else, continue checking next cards
				cardIndex := 0
				isCheckingCardsDone := false

				for !isCheckingCardsDone {
					card1 := strings.IndexByte(cardsOrder, sortedHands[handIndex].Cards[cardIndex])
					card2 := strings.IndexByte(cardsOrder, sortedHands[handIndex+1].Cards[cardIndex])
					if card1 < card2 {
						isCheckingCardsDone = true
					}

					if card1 > card2 {
						sortedHands[handIndex], sortedHands[handIndex+1] = sortedHands[handIndex+1], sortedHands[handIndex]
						isCheckingCardsDone = true
						isDone = false
						break
					}

					cardIndex++
				}
			}
			handIndex++
		}
	}

	return sortedHands
}

func part1() {
	start := time.Now()

	totalWinnings := 0
	part1Hands := sortHands(hands, "23456789TJQKA")

	// Calculate total winnings by summing hand's bid * rank
	for rank, hand := range part1Hands {
		totalWinnings += hand.Bid * (rank + 1)
	}

	fmt.Println("part1:", totalWinnings)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}

func part2() {
	start := time.Now()

	totalWinnings := 0
	part2Hands := sortHands(handsWithJoker, "J23456789TQKA")

	// Calculate total winnings by summing hand's bid * rank
	for rank, hand := range part2Hands {
		totalWinnings += hand.Bid * (rank + 1)
	}

	fmt.Println("\npart2:", totalWinnings)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}

func main() {
	part1()
	part2()
}
