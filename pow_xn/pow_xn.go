package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2.0000, -3))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return myPow(1/x, (-1 * n))
	}
	return x * myPow(x, n - 1)
}
