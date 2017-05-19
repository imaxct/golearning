package main

import (
	"fmt"
	"os"
)

var arr [200]int = func() (arr [200]int) {
	for i := range arr {
		arr[i] = i
	}
	return
}()

func main() {
	fmt.Println(arr[20])

	if f, err := os.Open("README.md"); err == nil {
		defer f.Close()
		defer fmt.Println("done")

		//buf := make([]byte, 10)
		var buf [200]byte
		n, r := f.Read(buf[0:20])
		if r == nil {
			fmt.Println(n)
			if n > 0 {
				ss := string(buf)
				fmt.Println(ss)
			} else {
				fmt.Println("empty file")
			}
		} else {
			fmt.Println(r.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
