package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var a, b float64
	fmt.Scanln(&a, &b)
	r := Rectangle{a, b}
	fmt.Println(r.Area())
}
