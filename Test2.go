package main

import (
	"fmt"
	"os"
)

/* var */

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var s string
	s = "xxxx"
	fmt.Println(s)

	var a, b, c, d int
	a = 10
	b = 20
	c = 30
	d = 40
	fmt.Println(a, b, c, d)

	var f, e = true, "error"
	fmt.Println(f, e)

	strs := []string{"a", "b"}
	fmt.Println(strs)

	for xx := range strs {
		fmt.Println(strs[xx])
	}


	fmt.Println(gcd(10, 15))

	file, err := os.Open("README.md")
	if err == nil {
		fmt.Println("success!")
		file.Close()
	}
}
