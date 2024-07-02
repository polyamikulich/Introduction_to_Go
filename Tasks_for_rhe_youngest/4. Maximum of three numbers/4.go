package main

import (
	"fmt"
	"math"
)

func max (x, y, z int) int {
	ans1 := math.Max(float64(x), float64(y))
	ans2 := math.Max(float64(ans1), float64(z))
	return int(ans2)
}

func main() {
	var x, y, z int
	fmt.Scanln(&x, &y, &z)
	fmt.Println(max(x, y, z))
}
