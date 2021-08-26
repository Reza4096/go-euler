package main

import "fmt"

func lcm(a int, b int) int {
	if a < b {
		a, b = b, a
	}

	for i, c := 2, a; a%b != 0; i++ {
		a = c * i
	}

	return a
}

func main() {
	ans := lcm(1, 2)

	for i := 3; i < 21; i++ {
		ans = lcm(ans, i)
	}

	fmt.Println(ans)
}
