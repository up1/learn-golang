package main

import "fmt"

func main() {
	p := make([]int, 10, 100)
	fmt.Println(p)
	fmt.Println(p==nil)
	fmt.Println(len(p))

	q := new([]int)
	fmt.Println(*q)
	fmt.Printf("Value = nill ? =>%t\n",*q==nil)
}
