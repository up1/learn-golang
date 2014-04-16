package main

import "fmt"

const (
	a0 = iota 
	a1 = iota 
	a2 = iota 
)

const (
	b0 = iota 
	b1 = iota 
	b2 = iota 
)

func main(){
	fmt.Println(a0)
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println(b0)
	fmt.Println(b1)
	fmt.Println(b2)
}