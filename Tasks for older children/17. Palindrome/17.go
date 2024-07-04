package main

import "fmt"

func palindrom(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	a := "Hello, world1 t 1dlrow ,olleH"
	fmt.Println(palindrom(a))
}
