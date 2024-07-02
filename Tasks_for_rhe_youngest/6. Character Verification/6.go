package main

import (
	"fmt"
)

func is_vowel(n string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "y", "A", "E", "I", "O", "U", "Y"}
	for _, v := range vowels {
		if n == v {
			return true
		}
	}
	return false
}

func main() {
	var n string
	fmt.Scanln(&n)
	fmt.Println(is_vowel(n))
}
