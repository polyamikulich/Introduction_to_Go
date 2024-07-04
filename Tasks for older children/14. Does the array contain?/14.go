package main

import (
	"fmt"
	"reflect"
)

func is(a []interface{}, elem interface{}) bool {
	elemType := reflect.TypeOf(elem)

	for v := range a {
		if elemType == reflect.TypeOf(v) {
			if elem == v {
				return true
			}
		}
	}
	return false
}

func main() {
	a := []interface{}{}
	var x interface{}

	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var input interface{}
		fmt.Scanln(&input)
		a = append(a, input)
	}

	fmt.Println("Enter element to search:")
	fmt.Scanln(&x)

	fmt.Println(is(a, x))
}
