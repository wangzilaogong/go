package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("go")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(" osx", os)
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("others", os)
	}
}
