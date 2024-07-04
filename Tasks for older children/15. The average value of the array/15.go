package main

import "fmt"

func middle(arr []int) float64 {
	var answer float64 = 0
	var n float64 = 0
	for _, v := range arr {
		answer += float64(v)
		n += 1
	}

	answer = answer / n
	return answer
}

func main() {
	arr := []int{1, 4, 2, 4, 5}
	fmt.Println(middle(arr))
}
