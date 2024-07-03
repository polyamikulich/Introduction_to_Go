package main

import "fmt"

func toFarenheit(c float64) float64 {
	c *= 1.8
	c += 32
	return c
}

func main() {
	var c float64
	fmt.Scanln(&c)
	fmt.Println(toFarenheit(c))
}
