package main

import (
	"fmt"
	"math"
)

func min_max(arr []int) [2]float64 {
	var min, max float64
	min = float64(arr[0])
	max = float64(arr[0])
	for _, v := range arr {
		min = math.Min(min, float64(v))
		max = math.Max(max, float64(v))
	}

	var answer [2]float64
	answer[0] = min
	answer[1] = max
	return answer
}

func main() {
	a := []int{1, 1}
	ans_min := min_max(a)[0]
	ans_max := min_max(a)[1]
	fmt.Println(ans_min, ans_max)
}
