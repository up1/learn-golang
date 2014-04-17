package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 1
	x[1] = 2
	x[2] = 3

	for i := 0; i < 5; i++ {
		fmt.Println(x[i])
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	for _, value := range x {
		fmt.Println(value)
	}
}
