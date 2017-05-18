package main

import "fmt"

const (
	c1 int64   = 20
	c2 float64 = 200.50
)

type NewVar float64

func (v NewVar) String() string {
	fmt.Printf("new value is %.4f\n", v)
	return fmt.Sprintf("new value is %.4f", v)
}

var pc [256]byte

func Init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	var sum int = 0
	var i uint
	for i = 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func main() {
	fmt.Println(c1, c2)
	var b float64 = 100.555
	var d int64 = int64(b)
	fmt.Println(b, d)
	var t NewVar = 55.555
	fmt.Println(t.String())

	Init()
	fmt.Println(Popcount(10))
}
