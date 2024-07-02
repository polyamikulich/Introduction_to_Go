package main

import (
	"fmt"
)

func fact (x int) int {
	var i, ans int = 1, 1
	
	for i <= x {
		ans *= i
		i++
	}

	return ans
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fact(n))
}
