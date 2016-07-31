package main

import "fmt"

const MAX_FIB = 40

var fibs [MAX_FIB]int

func fill_fibs() {
	fibs[0], fibs[1] = 0, 1
	for i := 2; i < MAX_FIB; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
}

func Zeckendorf(n int) []int {
	zeck := []int{}
	for i := MAX_FIB - 1; n > 0; i-- {
		if n >= fibs[i] {
			zeck = append(zeck, i)
			n -= fibs[i]
		}
	}
	return zeck
}

func main() {
	fill_fibs()
	for i := 0; i < MAX_FIB; i++ {
		fmt.Printf("%d\n", fibs[i])
	}
	fmt.Printf("Hello Brian. Zeckendorf(37)=%v\n", Zeckendorf(39))
	var right [100]int
	var sofar [100]int
	for i := 0; i < 100; i++ {
		sofar[i] = 0
	}
	sofar[0] = 0
	for i := 0; i < 50; i++ {
		right[i] = i - sofar[i]
		for j := i + 1; j <= i+right[i]; j++ {
			sofar[j]++
		}
	}
	for i := 0; i < 40; i++ {
		fmt.Printf("%2d: %v, %v, %v\n", i, Zeckendorf(right[i]), Zeckendorf(i), Zeckendorf(i+right[i]))
	}
}
