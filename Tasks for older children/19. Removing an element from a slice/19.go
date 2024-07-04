package main

import "fmt"

func delete(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	a := []int{1, 5, 3, 8}
	ans := delete(a, 2)
	for _, v := range ans {
		fmt.Println(v)
	}
}
