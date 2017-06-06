package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

/* just basic */
func main() {
	s := md5.New()
	x := s.Sum([]byte("abcd"))
	fmt.Println(hex.EncodeToString(x))
}
