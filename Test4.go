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
		buf := make([]byte, 10)
		n, r := f.Read(buf);
		if (r == nil){
			fmt.Println(n)
			if n > 0 {
				ss := string(buf)
				fmt.Println(ss)
			}else {
				fmt.Println("empty file")
			}
		}else {
			fmt.Println(r.Error())
		}
	}else {
		fmt.Println(err.Error())
	}
}