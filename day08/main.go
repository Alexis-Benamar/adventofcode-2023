package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Node struct {
	L string
	R string
}

//go:embed input.txt
var data string
var dataAsBlocks, lines, nodes []string
var directions string
var nodeMap map[string]Node

func init() {
	data = strings.TrimRight(data, "\n")
	dataAsBlocks = strings.Split(data, "\n\n\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(dataAsBlocks[0], "\n")
	directions = lines[0]

	nodeMap = make(map[string]Node, len(lines[2:]))
	for _, line := range lines[2:] {
		nodeMap[line[:3]] = Node{
			L: line[7:10],
			R: line[12:15],
		}
	}
}

func getNextNodeIndex(node Node, direction byte) string {
	if string(direction) == "L" {
		return node.L
	} else {
		return node.R
	}
}

/*
In part 1, we just follow the nodes until we land on node ZZZ
*/
func part1() {
	if match, _ := regexp.MatchString("A{3}", dataAsBlocks[0]); !match {
		fmt.Printf("part1: cannot run (no starting node AAA found)\n\n")
		return
	}

	start := time.Now()

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
		nextNodeIndex = getNextNodeIndex(currentNode, direction)

		// Stop when landing on node ZZZ
		if nextNodeIndex == "ZZZ" {
			keepChecking = false
		}

		currentNode = nodeMap[nextNodeIndex]
		dIndex++
	}

	fmt.Println("part1:", dIndex)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n\n", elapsed.Seconds())
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

/*
In part 2, we can't do the whole path for each ghost nodes, it will take a huge nb of steps (tens of billions)
So instead, we calculate each ghost node's step number, and find the Least Common Multiple between them,
As it corresponds to the nb of step required for all of them to end in ..Z
*/
func part2() {
	start := time.Now()

	reNodeEndingWithA := regexp.MustCompile(`\w{2}A`)
	ghostNodes := reNodeEndingWithA.FindAllString(dataAsBlocks[0], -1)
	stepsPerNode := make([]int, len(ghostNodes))

	// Get nb of steps for each ghost node
	for nodeI, node := range ghostNodes {
		nextNodeIndex := node
		keepChecking := true
		dIndex := 0

		for keepChecking {
			// The modulo will always return an index between 0 (included) & len(directions) (not included)
			direction := directions[dIndex%len(directions)]

			// Pick next node depending on direction
			nextNodeIndex = getNextNodeIndex(nodeMap[nextNodeIndex], direction)

			if match, _ := regexp.MatchString(`\w{2}Z`, nextNodeIndex); match {
				keepChecking = false
			}

			dIndex++
		}

		stepsPerNode[nodeI] = dIndex
	}

	fmt.Println("part2:", LCM(1, 1, stepsPerNode...))

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}

func main() {
	part1()
	part2()
}
