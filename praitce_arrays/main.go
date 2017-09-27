package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// An array has a fixed size called 'slice', on the other hand,
	// is a dynamiclly-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.
	var slice []int = primes[1:4]
	fmt.Println(slice)

	names := [4]string{
		"John",
		"Roy",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	b := names[1:3]
	fmt.Println(c, b)

	b[0] = "XXX"
	fmt.Println(c, b)
	fmt.Println(names)

	q := []int{2, 3, 54, 5}
	fmt.Println(q)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)

	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	// Slice the slice to give it zero length.
	numbers = numbers[:0]
	printSlice(numbers)

	// extend its length
	numbers = numbers[:4]
	printSlice(numbers)

	// Drop its first two values
	numbers = numbers[2:]
	printSlice(numbers)

	var nil_slice []int
	printSlice(nil_slice)
	if nil_slice == nil {
		fmt.Println("nil!")
	}

	// use build-in make function to create dynamically-sized arrays
	ex1 := make([]int, 5)
	printSlice(ex1)
	// make(slice type, capacity, length)
	ex2 := make([]int, 0, 5)
	printSlice(ex2)

	ex3 := ex2[:2]
	printSlice(ex3)

	ex4 := ex3[2:5]
	printSlice(ex4)

	ex5 := append(ex4, 2, 3, 4, 5)
	printSlice(ex5)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}


	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
