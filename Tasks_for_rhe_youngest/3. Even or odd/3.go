package main

import (
	"fmt"
)

func is (x int) bool {
	if x % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println(is(x))
}
