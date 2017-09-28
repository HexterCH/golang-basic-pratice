package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// a method is just a function with a receiver argument.
// (v Vertex) is receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// methods with pointer receivers.
// methods with pointer receivers can modify the value to which the receiver points
// receive is pointer or not is not important
// golang convenience, equal to (&v).Scale(10)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// match pointer type use &, transform to pointer use *

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v,10)
	fmt.Println(Abs(v))

	example := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", example, example.Abs())
	example.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", example, example.Abs())
}