package main

import "fmt"

func goTest(ch chan int) {
	for i := 0; i<20; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 20)
	go goTest(ch)
	for x := range ch {
		fmt.Println(x)
	}
}
