package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed example.txt
var data string

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}
}

func main() {
	start := time.Now()

	// Code here

	elapsed := time.Since(start)
	fmt.Println("\nExecution time", elapsed.Seconds(), "s")
}
