package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

type Crawler struct {
	pos, lastPos int
}

type Point struct {
	x, y int
}

//go:embed example.txt
var data string

var grid string
var blocks, lines []string
var nbCol, nbRow, nbPointInside, startPos, steps int
var path, polygon []int
var dirToIndex map[string]int
var canMoveTo = map[string]string{
	"up": "7|F",
  "right": "J-7",
  "down": "J|L",
  "left": "L-F",
}
var connectingPipes = map[string][]string{
  "F": {"right", "down"},
  "|": {"up", "down"},
  "L": {"up", "right"},
  "-": {"left", "right"},
  "J": {"left", "up"},
  "7": {"left", "down"},
}

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	blocks = strings.Split(data, "\n\n")
	lines = strings.Split(blocks[0], "\n")
	grid = strings.Join(lines, "")

	fmt.Println(blocks[0])

	nbRow = len(lines)
	nbCol = len(lines[0])

	dirToIndex = map[string]int{
		"up": -nbCol,
		"right": +1,
		"down": nbCol,
		"left": -1,
	}
}

func getChar(i int) string {
	return fmt.Sprintf("%c", grid[i])
}

func moveCrawler(crawler *Crawler) {
	currentPipe := getChar(crawler.pos)

	if currentPipe == "S" {
		return
	}

	possibleDirs := connectingPipes[currentPipe]

	if crawler.pos + dirToIndex[possibleDirs[0]] == crawler.lastPos {
		crawler.lastPos = crawler.pos
		crawler.pos += dirToIndex[possibleDirs[1]]
	} else {
		crawler.lastPos = crawler.pos
		crawler.pos += dirToIndex[possibleDirs[0]]
	}
}

func possiblyAddVertexToPolygon(crawler *Crawler) {
	if getChar(crawler.pos) != "-" && getChar(crawler.pos) != "|" {
		polygon = append(polygon, crawler.pos)
	}
}

func getXY(point int) (x, y float64) {
	return float64(point % nbCol), float64((point % nbCol) % nbRow)
}

func checkIfPointIsInside(point int) bool {
	x, y := getXY(point)
	inside := false

	var p1x, p1y, p2x, p2y float64
	p1x, p1y = getXY(polygon[0])

	for i := 1; i < len(polygon); i++ {
		p2x, p2y = getXY(polygon[i % len(polygon)])

		if y > math.Min(p1y, p2y) {
			if y <= math.Max(p1y, p2y) {
				if x <= math.Max(p1x, p2x) {
					xIntersect := (y - p1y) * (p2x - p1x) / (p2y - p1y) + p1x

					if p1x == p2x || x <= xIntersect {
						inside = !inside
					}
				}
			}
		}

		p1x, p1y = p2x, p2y
	}

	return inside
}

func main() {
	// part 1
	startPart1 := time.Now()

	// Get start position "S", and place crawler here
	startPos = strings.Index(grid, "S")
	crawler := Crawler{
		pos: startPos,
		lastPos: startPos,
	}

	path = append(path, startPos)
	polygon = append(polygon, startPos)

	var possibleDirFromStart []int
	for dir, possibleChar := range canMoveTo {
		if startPos + dirToIndex[dir] < 0 || startPos + dirToIndex[dir] > len(grid) {
			continue
		}

		charToCheck := getChar(startPos + dirToIndex[dir])
		isConnected := strings.Index(possibleChar, charToCheck) != -1

		if isConnected {
			possibleDirFromStart = append(possibleDirFromStart, startPos + dirToIndex[dir])
		}
	}

	crawler.pos = possibleDirFromStart[0]
	steps++
	possiblyAddVertexToPolygon(&crawler)

	for true {
		path = append(path, crawler.pos)
		moveCrawler(&crawler)

		steps++
		if getChar(crawler.pos) == "S" {
			break
		}

		possiblyAddVertexToPolygon(&crawler)
	}

	fmt.Println("part1:", steps / 2)

	elapsedPart1 := time.Since(startPart1)
	fmt.Printf("Execution time %f s\n", elapsedPart1.Seconds())

	// Part 2
	startPart2 := time.Now()

	// for each line
	for i := 0; i < nbRow; i++ {
		// get left & right boundary
		var leftB, rightB int
		for j := i * nbCol; j < i * nbCol + nbCol; j++ {
			if slices.Contains(path, j) {
				leftB = j
				break
			}
		}
		for j := i * nbCol + (nbCol - 1); j > i * nbCol; j-- {
			if slices.Contains(path, j) {
				rightB = j
				break
			}
		}
		fmt.Println("\nrow", i, "leftB", leftB, "rightB", rightB)

		// check from leftB to rightB, if point is inside polygon
		for pointIdx := leftB; pointIdx <= rightB; pointIdx++ {
			// ignore points on path
			if slices.Contains(path, pointIdx) {
				fmt.Println("point", pointIdx, "ignored")
				continue
			}

			// todo: raycast
			if checkIfPointIsInside(pointIdx) {
				fmt.Println("point", pointIdx, "is inside !")
				nbPointInside++
			} else {
				fmt.Println("point", pointIdx, "outside :(")
			}
		}
	}

	fmt.Println("\npart2:", nbPointInside)

	elapsedPart2 := time.Since(startPart2)
	fmt.Printf("Execution time %f s\n", elapsedPart2.Seconds())
}
