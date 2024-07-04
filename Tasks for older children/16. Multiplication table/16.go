package main

import "fmt"

func Print(n float64, k int) {
	for i := 0; i <= k; i++ {
		var p float64 = float64(i) * n
		fmt.Printf("%v * %v = %v\n", n, i, p)
	}
}

func main() {
	var n float64
	var k int //к - число, до которого прописать таблицу
	fmt.Scanln(&n, &k)
	Print(n, k)
}
