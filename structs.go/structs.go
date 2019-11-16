package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// 结构体指针
	v := Vertex{2, 3}
	fmt.Println(v.X)
	fmt.Println(&v)
}
