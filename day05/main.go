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

//go:embed example.txt
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
	seeds := make([]int, len(seedNbAsStr))
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

func getSeedLocation(seed int) int {
	/*
		soil := getMapped(seed, seedToSoil)
		fertilizer := getMapped(soil, soilToFertilizer)
		water := getMapped(fertilizer, fertilizerToWater)
		light := getMapped(water, waterToLight)
		temp := getMapped(light, lightToTemp)
		humidity := getMapped(temp, tempToHumidity)
		location := getMapped(humidity, humidityToLocation)

		return location
	*/

	return 0
}

func main() {
	start := time.Now()

	// Code here
	// TODO: remove, test data
	fmt.Println(mappings)

	res := getMapped(98, mappings["seed-to-soil"])

	fmt.Println("returned", res)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time %f s\n", elapsed.Seconds())
}
