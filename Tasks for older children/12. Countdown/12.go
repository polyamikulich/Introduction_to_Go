package main

import "fmt"

func Countdown(n int) {
	for n > 0 {
		fmt.Println(n)
		n -= 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	Countdown(n)
}
