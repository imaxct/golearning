package main

import "fmt"

func main() {
	for i, ch := range "哈哈哈\x80" {
		fmt.Printf("%#U %d\n", ch, i)
	}

	switch 1 + 2 {
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("none")
	}

	var t interface{}
	t = "aaa"
	switch t.(type) {
	default:
		fmt.Printf("unknow type %T", t)
	case string:
		fmt.Printf("string\n")
	case int:
		fmt.Printf("int\n")
	}
}
