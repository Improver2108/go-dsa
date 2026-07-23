package main

import (
	"fmt"
	"time"

	slidingwindow "github.com/improver2108/godsa/sliding_window"
)

func main() {
	start := time.Now()
	res := slidingwindow.RunTotalFruit()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
