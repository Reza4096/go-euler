package main

import (
	"fmt"
	"os"
)

func main() {
	for a := 2; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				os.Exit(0)
			}
		}
	}
}
