package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	fibo := 0
	sum := 0
	for fibo < 4000000 {
		fibo = a + b
		a = b
		b = fibo
		if fibo%2 == 0 {
			sum += fibo
		}
	}
	fmt.Println(sum)
}
