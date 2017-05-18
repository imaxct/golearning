package main

import "fmt"

/* var */

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
}
