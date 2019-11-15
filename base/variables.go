package main

import "fmt"

// 函数外
var a, b, c int = 1, 2, 3

func main() {
	// 函数内
	a := 2
	var i int
	fmt.Println(a, b, c, i)
}
