package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	// Switch evaluation order
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		// if os is OS X. These lines will never be called
		fmt.Println("hello, i am here")
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
