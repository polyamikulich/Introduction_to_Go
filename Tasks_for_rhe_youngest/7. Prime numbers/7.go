package main

import (
	"fmt"
)

func is_prime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 2; i <= n; i++ {
		if is_prime(i) {
			fmt.Println(i)
		}
	}
}
