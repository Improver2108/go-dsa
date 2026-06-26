package main

import (
	"fmt"
	"time"

	"github.com/improver2108/godsa/dp1d"
)

func main() {
	start := time.Now()
	res := dp1d.RunLengthOfLIS()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
