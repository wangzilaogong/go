package main

import "os"
import "fmt"

func main() {
	input, err :=  os.Open("input.dat")
	if err !== nil {
		fmt.Println('WRONG')
		return
	}
	defer input.Close()


}
