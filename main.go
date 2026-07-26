package main

import (
	"fmt"
	"time"

	"github.com/improver2108/godsa/dp2d"
)

func main() {
	start := time.Now()
	res := dp2d.RunLongestCommonSubsequence()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
