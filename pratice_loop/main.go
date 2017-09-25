package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 1

	// for is Go's while
	for sum < 1000 {
		sum += sum
	}


	fmt.Println(sum)
	//fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 4, 20),
	)
	z := float64(2)
	fmt.Println(Sqrt(z))
	fmt.Println(math.Sqrt(z))
}

//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}

func Sqrt(x float64) float64 {
	res := 1.0
	for n := 0; n < 10; n++ {
		res = res - ((res * res - x) / (2 * res))
	}
	return res
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}