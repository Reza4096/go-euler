package main

import "fmt"

func isPalindrome(n int) bool {
	a, m := 0, n
	for m != 0 {
		a = a*10 + m%10
		m /= 10
	}

	if a == n {
		return true
	}

	return false

}

func main() {
	ans, k := 0, 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			k = i * j
			if isPalindrome(k) && k > ans {
				ans = k
			}
		}
	}

	fmt.Println(ans)
}
