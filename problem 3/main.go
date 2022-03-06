package main

import (
	"fmt"
)

func main() {
	n := 600851475143
	var max int

	if n%2 == 0 {
		max = 2
	}
	for n%2 == 0 {
		n = n / 2
	}

	for j := 3; j*j <= n; j += 2 {
		for n%j == 0 {
			n = n / j
		}
	}

	if n > 2 {
		max = n
	}

	fmt.Println(max)
}
