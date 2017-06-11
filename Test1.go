package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init b")
}

/* just basic */
func main() {
	s := md5.New()
	x := s.Sum([]byte("abcd"))
	fmt.Println(hex.EncodeToString(x))
}
