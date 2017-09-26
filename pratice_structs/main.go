package main

import (
	"fmt"
)

// a collection of fields
type Vertex struct {
	X, Y int
}

var	(
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &v1
)

func main() {
	v := Vertex{1, 2}
	// struct fields are accessed useing a dot
	v.X = 4
	pointer_value := &v
	pointer_value.Y = 1e9
	fmt.Println(v)
	fmt.Println(pointer_value)
	fmt.Println(v1, p, v2, v3)
}
