package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var example interface{}
	describe(example)

	example = 42
	describe(example)

	example = "TT man"
	describe(example)

	var i I

	// In some languages this would trigger a null pointer exception,
	// but in Go it is common to write methods that gracefully handle being called with a nil receiver
	// (as with the method M in this example.)
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello	"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}