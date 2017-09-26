package main

import "fmt"

func main() {
	i, j := 42, 277

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	// p = 21, if p is a pointer to another value, can't be done like this
	fmt.Println(p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
