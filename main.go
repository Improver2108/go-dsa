package main

import (
	"fmt"
	"time"

	"github.com/improver2108/godsa/graph"
)

func main() {
	start := time.Now()
	res := graph.RunIslandsAndTreasure()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
