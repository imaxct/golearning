package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"xx": {
		11.11, 22.22,
	},
	"yy": {
		33.33, 44.44,
	},
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Lat * v.Lat + v.Long * v.Long)
}

func main() {
	fmt.Println(m)
	ff :=m["xx"]
	fmt.Println(ff.Abs())
}
