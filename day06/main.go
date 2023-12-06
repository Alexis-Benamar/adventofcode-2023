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

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
}

func checkRace(times, distances []string) int {
	res := 1

	for raceIndex, timeStr := range times {
		time, _ := strconv.Atoi(timeStr)
		distanceToBeat, _ := strconv.Atoi(distances[raceIndex])
		validOptions := 0

		for i := 0; i <= time; i++ {
			distance := i * (time - i)

			if distance > distanceToBeat {
				validOptions++
			}

		}

		res *= validOptions
	}

	return res
}

func main() {

	// Code here
	re := regexp.MustCompile(`\d+`)

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	mergedTime := strings.Join(times, "")
	mergedDistance := strings.Join(distances, "")
	timeArray := []string{mergedTime}
	distanceArray := []string{mergedDistance}

	startPart1 := time.Now()
	part1 := checkRace(times, distances)
	elapsedPart1 := time.Since(startPart1)
	fmt.Println("part1:", part1)
	fmt.Printf("Execution time %f s\n", elapsedPart1.Seconds())

	startPart2 := time.Now()
	part2 := checkRace(timeArray, distanceArray)
	elapsedPart2 := time.Since(startPart2)

	fmt.Println("\npart1:", part2)
	fmt.Printf("Execution time %f s\n", elapsedPart2.Seconds())
}
