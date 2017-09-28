package main

import (
	"time"
	"fmt"
)

// Go programs express error state with error values.
// The error type is a built-in interface similar to fmt.Stringer:

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}