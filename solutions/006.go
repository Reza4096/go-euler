package main

import "fmt"

func main() {
	a, b := 0, 0
	for i := 1; i < 101; i++ {
		a += i
		b += i * i
	}

	fmt.Println(a*a - b)

}
