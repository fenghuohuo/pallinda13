package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := float64(1)
	delta := 1.0
	var oldRes float64
	i := 1
	for ; math.Abs(delta) > 0.000001; i++ {
		oldRes = res
		res = res - (math.Pow(res, 2)-x)/(2*res)
		delta = oldRes - res
	}
	fmt.Println(i)
	return res
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
