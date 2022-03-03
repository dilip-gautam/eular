package main

import "fmt"

func main() {
out:
	for a := 1; a < 1_000; a++ {
		for b := 1; b < 1_000; b++ {
			for c := 1; c < 1_000; c++ {
				if ((a*a)+(b*b) == (c * c)) && (a+b+c == 1000) && ((a < b) && (b < c)) {
					product := a * b * c
					fmt.Println(product)
					break out
				}
			}
		}
	}
}
