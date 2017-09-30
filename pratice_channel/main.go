package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s	{
		sum += v
	}

	c <- sum // send sum to c channel
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	fmt.Println(s[:len(s)/2])
	go sum(s[:len(s)/2], c)

	fmt.Println(s[len(s)/2:])
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// channels can be bufferd.
	// Provide the buffer length as the second argument to make to initialize a buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	// when the buffer channel is full, will raise deadlock and the process in channels will be set to asleep
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fibonacci_channel := make(chan int, 10)
	go fibonacci(cap(fibonacci_channel), fibonacci_channel)
	for i := range fibonacci_channel {
		fmt.Println(i)
	}

	// will raise panic: send on closed channel
	// fibonacci_channel <- 1
}
