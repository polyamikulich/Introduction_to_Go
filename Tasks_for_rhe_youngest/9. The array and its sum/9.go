package main

import "fmt"

func sum(array []int) int {
	answer := 0
	for _, v := range array {
		answer += v
	}

	return answer
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(array))
}