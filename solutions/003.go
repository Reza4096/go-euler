package main

import "fmt"

func main() {
	a, p := 600851475143, 1

	for i := 3; a != 1; i += 2 {
		if a%i == 0 {
			p = i
			for a%i == 0 {
				a /= i
			}
		}
	}

	fmt.Println(p)
}
