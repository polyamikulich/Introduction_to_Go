package main

import (
	"fmt"
)

func per(n string) string {
	answer := ""
	for i := 0; i < len(n); i++ {
		answer += string(n[len(n) - 1 - i])
	}

	return answer
}

func main() {
	var n string
	n = "hello, world, my name is Polina"
	fmt.Print(per(n))
}
