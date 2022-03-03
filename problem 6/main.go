package main

import (
	"fmt"
)

func main() {
	n := 100
	sum := n * (n + 1) / 2
	squareOfSum := sum * sum
	sumOfSquare := n * (2*n + 1) * (n + 1) / 6
	fmt.Println(squareOfSum - sumOfSquare)
}
