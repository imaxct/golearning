package main

import "fmt"

func main() {
	h := map[string]int{
		"aaa": 10,
		"bb": 12,
	}
	fmt.Println(h["bb"])
	fmt.Println(h["cc"])
}