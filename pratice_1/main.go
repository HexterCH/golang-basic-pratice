package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// func name(params) (return result type) {
// code
// }
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(add(42, 13))
}
