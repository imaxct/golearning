package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	p := make([]int, 5)
	fmt.Println(p)
	p[0] = 1
	p[1] = 10
	fmt.Println(p)
	p = append(p, 6, 2, 3)
	fmt.Println(p)

	for k, v := range p {
		fmt.Println(k, v)
	}
}
