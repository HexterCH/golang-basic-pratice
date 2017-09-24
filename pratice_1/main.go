package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// variables
// package level
var c, python, java bool

const Pi = 3.14

// func name(params) (return result type) {
// code
// }
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return will returns the named return values
	// naked return
	return
}

func main() {
	// := use to replace the variable declaration like var xxx int = 1
	// only can use in function level
	k := 3
	fmt.Println(k)

	// variable function level
	var i int
	var z, j, h = true, false, "no!"
	fmt.Println(z, j, h)
	fmt.Println(i, c, python, java)
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(add(42, 13))
	fmt.Println(split(17))
	fmt.Println(Pi)
}
