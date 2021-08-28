package main

import (
	"fmt"
	"math"
)

func main() {
	var ans int

	for i, a, isPrime := 3, 1, true; a != 10001; i += 2 {
		for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			a += 1
			ans = i
		}
		isPrime = true
	}

	fmt.Println(ans)
}
