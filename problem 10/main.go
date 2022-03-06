package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	prime := true
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
func main() {
	sum := 0
	for i := 2; i <= 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}