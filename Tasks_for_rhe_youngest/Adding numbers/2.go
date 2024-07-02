package main

import (
	"fmt"
)

func add (x, y int) int {
	return x + y
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(add(x, y))
}
