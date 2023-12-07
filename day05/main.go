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
	Source      int
	Length      int
}

//go:embed input.txt
var data string
var blocks []string
var seeds []int
var seedsAsRange [][]int
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
				Source:      source,
				Length:      length,
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
		if nbToCheck >= potentialMapping.Source && nbToCheck <= potentialMapping.Source+potentialMapping.Length {
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
From a mapped number, gets its original number
*/
func getMappedReverse(nbToCheck int, mappingList []Mapping) int {
	// Search for matching mapping
	var mapping Mapping
	for _, potentialMapping := range mappingList {
		if nbToCheck >= potentialMapping.Destination && nbToCheck <= potentialMapping.Destination+potentialMapping.Length {
			mapping = potentialMapping
			break
		}
	}

	// If no mapping found, directly return nbToCheck
	if mapping == (Mapping{}) {
		return nbToCheck
	}

	// Else, get the original number
	if mapping.Source < mapping.Destination {
		return nbToCheck - (mapping.Destination - mapping.Source)
	} else {
		return nbToCheck + (mapping.Source - mapping.Destination)
	}
}

/*
Returns location, going through all mappings in order, for a given seed
*/
func getLocationFromSeed(seed int) int {
	soil := getMapped(seed, mappings["seed-to-soil"])
	fertilizer := getMapped(soil, mappings["soil-to-fertilizer"])
	water := getMapped(fertilizer, mappings["fertilizer-to-water"])
	light := getMapped(water, mappings["water-to-light"])
	temp := getMapped(light, mappings["light-to-temperature"])
	humidity := getMapped(temp, mappings["temperature-to-humidity"])
	location := getMapped(humidity, mappings["humidity-to-location"])

	return location
}

/*
Returns seed, going through all mappings in reverse order, for a given location
*/
func getSeedFromLocation(location int) int {
	humidity := getMappedReverse(location, mappings["humidity-to-location"])
	temp := getMappedReverse(humidity, mappings["temperature-to-humidity"])
	light := getMappedReverse(temp, mappings["light-to-temperature"])
	water := getMappedReverse(light, mappings["water-to-light"])
	fertilizer := getMappedReverse(water, mappings["fertilizer-to-water"])
	soil := getMappedReverse(fertilizer, mappings["soil-to-fertilizer"])
	seed := getMappedReverse(soil, mappings["seed-to-soil"])

	return seed
}

func getLowestLocationFromSeeds(seedNumbers []int) (lowestLocation int) {
	for _, seed := range seeds {
		location := getLocationFromSeed(seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return
}

func part1() {
	start := time.Now()

	// Check for all seeds's corresponding locations, and keep the lowest
	var lowestLocation int
	for _, seed := range seeds {
		location := getLocationFromSeed(seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("part1:", lowestLocation)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n\n", elapsed.Seconds())
}

func part2() {
	start := time.Now()

	// Get seed numbers as range (seedNbBase + length of range)
	numbersAsStr := reNum.FindAllString(blocks[0][7:], -1)
	for i := 0; i < len(numbersAsStr)-1; i += 2 {
		seedNbBase, _ := strconv.Atoi(numbersAsStr[i])
		seedNbRange, _ := strconv.Atoi(numbersAsStr[i+1])
		seedsAsRange = append(seedsAsRange, []int{seedNbBase, seedNbRange})
	}

	var location, seed int

	/*
	Check every location number's corresponding seed number
	Stop when finding a seed number that's inside one of the given seed ranges
	*/
	for true {
		seed = getSeedFromLocation(location)

		// Search for matching seed range
		var seedRange []int
		for _, potentialSeedRange := range seedsAsRange {
			if seed >= potentialSeedRange[0] && seed <= potentialSeedRange[0] + potentialSeedRange[1] {
				seedRange = potentialSeedRange
				break
			}
		}

		// If no seed range found, continue checking for next location
		if len(seedRange) == 0 {
			location += 1
			continue
		}

		break
	}

	fmt.Printf("part2: %d (seed number %d)\n", location, seed)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}

func main() {
	part1()
	part2()
}
