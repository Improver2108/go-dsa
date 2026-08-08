package main

import (
	"fmt"
	"time"

	"github.com/improver2108/godsa/dp2d"
)

func main() {
	start := time.Now()
	res := dp2d.RunRegularExpression()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println("Time taken:", elapsed)
}
