package main

import (
	"time"
	"fmt"
)

func say(s string, c chan int){

	for i := 0; i<=10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	c <- 0
}

func main() {
	c := make(chan int)
	go say("hhh", c)
	x := <-c
	fmt.Println(x)
}
