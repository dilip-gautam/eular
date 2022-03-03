package main

import "fmt"

func main() {
	numberLongestChain := 1
	chainLen := 1
	for i := 2; i <= 1000000; i++ {
		count := 1
		start := i
		for start != 1 {
			if start%2 == 0 {
				start /= 2
			} else {
				start = start*3 + 1
			}
			count++
		}
		if count > chainLen {
			chainLen = count
			numberLongestChain = i
		}
	}
	fmt.Println(numberLongestChain)

}
