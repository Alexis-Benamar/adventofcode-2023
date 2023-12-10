package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Node struct {
	L string
	R string
}

//go:embed input.txt
var data string
var lines, nodes []string
var directions string
var nodeMap map[string]Node

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
	directions = lines[0]

	nodeMap = make(map[string]Node, len(lines[2:]))
	for _, line := range lines[2:] {
		nodeMap[line[:3]] = Node{
			L: line[7:10],
			R: line[12:15],
		}
	}
}

func main() {
	start := time.Now()

	// Code here
	keepChecking := true
	nextNodeIndex := "AAA"
	currentNode := nodeMap[nextNodeIndex]
	dIndex := 0

	for keepChecking {
		// The modulo will always return an index between 0 (included) & len(directions) (not included)
		direction := directions[dIndex%len(directions)]

		// Debug
		// fmt.Println(nextNodeIndex, currentNode, string(direction))

		// Pick next node depending on direction
		if string(direction) == "L" {
			nextNodeIndex = currentNode.L
		} else {
			nextNodeIndex = currentNode.R
		}

		// Stop when landing on node ZZZ
		if nextNodeIndex == "ZZZ" {
			keepChecking = false
		}

		currentNode = nodeMap[nextNodeIndex]
		dIndex++
	}

	fmt.Println("part1:", dIndex)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}
