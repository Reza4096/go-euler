package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := 2

	for i := 3; i < 2000000; i += 2 {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			sum += i
		}
	}

	fmt.Println(sum)
}
