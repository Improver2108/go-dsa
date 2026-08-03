package main

import (
	"fmt"
	"time"

	"github.com/improver2108/godsa/hashing"
)

func main() {
	start := time.Now()
	res := hashing.RunCountValidPrefixes()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
