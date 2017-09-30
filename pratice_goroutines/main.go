package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go function_name set a goroutine thread
	go say("world")
	// starts a new goroutine running
	say("hello")

	// Goroutines run in the same address space, so access to shared memory must be synchronized
}
