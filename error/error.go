package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	When time.Time
	What string
}

func run() {
	return &Myerror{
		time.Now(),
		"itesss",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)

	}
}
