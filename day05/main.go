package main

import (
	_ "embed"
	"fmt"
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
var lines []string

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
}

func getMapped(nbToCheck int, mapToCompare [][]int) int {
	/*
		mapping = find mapping in mapToCompare where source < nbToCheck < (source + length)
		if (mapping not found)
			return nbToCheck
		else
			if (destination < source)
				return seedNb - (source - destination)
			else
				return seedNb + (destination - source)
	*/
	var mapping []int

	for _, potentialMap := range mapToCompare {
		if nbToCheck >= potentialMap[1] && nbToCheck <= potentialMap[1] + potentialMap[2] {
			mapping = potentialMap
			break
		}
	}

	fmt.Println(nbToCheck, mapping)

	if len(mapping) == 0 {
		return nbToCheck
	}

	if mapping[0] < mapping[1] {
		return nbToCheck - (mapping[1] - mapping[0])
	} else {
		return nbToCheck + (mapping[0] - mapping[1])
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
	res := getMapped(6, [][]int{{0, 5, 4}, {6, 10, 4}, {52, 50, 48}})

	fmt.Println("returned", res)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time %f s\n", elapsed.Seconds())
}
