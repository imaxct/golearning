package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	done := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick..")
		case <-done:
			fmt.Println("done..")
			return
		default:
			fmt.Println(".....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}