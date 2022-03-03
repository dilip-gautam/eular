package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1
	triangle := 0
	for {
		triangle += n
		noDivisors := 0
		to := int(math.Sqrt(float64(triangle)))
		for i := 1; i <= to; i++ {
			if triangle%i == 0 {
				noDivisors += 2 //every divisor upto square root there is corresponding divisor above square root
			}
		}
		if noDivisors >= 500 {
			fmt.Println(triangle)
			break
		}
		n++
	}
}
