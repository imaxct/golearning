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

func main() {
	fmt.Println(c1, c2)
	var b float64 = 100.555
	var d int64 = int64(b)
	fmt.Println(b, d)
	var t NewVar = 55.555
	fmt.Println(t.String())
}
