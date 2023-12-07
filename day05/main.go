package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Mapping struct {
	Destination int
	Source int
	Length int
}

//go:embed input.txt
var data string
var blocks []string
var seeds []int
var mappings map[string][]Mapping
var reNum = regexp.MustCompile(`\d+`)

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	blocks = strings.Split(data, "\n\n")

	// Get seed numbers from first block
	seedNbAsStr := reNum.FindAllString(blocks[0][7:], -1)
	seeds = make([]int, len(seedNbAsStr))
	for i, seedNb := range seedNbAsStr {
		seeds[i], _ = strconv.Atoi(seedNb)
	}

	// Add remaining as []Mapping in a map (in this house, we love maps)
	mappings = make(map[string][]Mapping, len(blocks[1:]))
	for _, block := range blocks[1:] {
		blockLines := strings.Split(block, "\n")
		mappingTitle := strings.Split(blockLines[0], " ")[0]

		for _, line := range blockLines[1:] {
			lineNbs := strings.Split(line, " ")
			destination, _ := strconv.Atoi(lineNbs[0])
			source, _ := strconv.Atoi(lineNbs[1])
			length, _ := strconv.Atoi(lineNbs[2])

			mappings[mappingTitle] = append(mappings[mappingTitle], Mapping{
				Destination: destination,
				Source: source,
				Length: length,
			})
		}
	}
}

/*
	Compares a number against an array of mappings.
	If number correspond to a certain mapping, returns the mapped number.
	Else, returns number directly.
*/
func getMapped(nbToCheck int, mappingList []Mapping) int {
	// Search for matching mapping
	var mapping Mapping
	for _, potentialMapping := range mappingList {
		if nbToCheck >= potentialMapping.Source && nbToCheck <= potentialMapping.Source + potentialMapping.Length {
			mapping = potentialMapping
			break
		}
	}

	// If no mapping found, directly return nbToCheck
	if mapping == (Mapping{}) {
		return nbToCheck
	}

	// Else, get the mapped number
	if mapping.Destination < mapping.Source {
		return nbToCheck - (mapping.Source - mapping.Destination)
	} else {
		return nbToCheck + (mapping.Destination - mapping.Source)
	}
}

/*
	Returns location, going through all mappings in order, for a given seed
*/
func getSeedLocation(seed int) int {
	soil := getMapped(seed, mappings["seed-to-soil"])
	fertilizer := getMapped(soil, mappings["soil-to-fertilizer"])
	water := getMapped(fertilizer, mappings["fertilizer-to-water"])
	light := getMapped(water, mappings["water-to-light"])
	temp := getMapped(light, mappings["light-to-temperature"])
	humidity := getMapped(temp, mappings["temperature-to-humidity"])
	location := getMapped(humidity, mappings["humidity-to-location"])

	return location
}

// TODO
func getLowestLocationFromSeeds(seedNumbers []int) {
	/*
		Actual content of main
		Use for both part1 and part2, just give it an array of seed numbers
	*/
}

func main() {
	start := time.Now()

	// Check for all seeds's corresponding locations, and keep the lowest
	var lowestLocation int
	for _, seed := range seeds {
		location := getSeedLocation(seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("part1:", lowestLocation)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time %f s\n", elapsed.Seconds())
}
