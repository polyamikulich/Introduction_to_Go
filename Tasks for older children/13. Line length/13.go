package main

import "fmt"

func len(str string) int {
	var result int = 0
	for range str {
		result += 1
	}

	return result
}

func main() {
	my_string := "Hello World"
	fmt.Println(len(my_string))
}
