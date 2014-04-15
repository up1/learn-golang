package main

import "fmt"

func show(input []string) {
	fmt.Println(input)
	fmt.Printf("Length=%d\n", len(input))
	fmt.Printf("Capacity=%d\n", cap(input))	
}

func main() {
	letters := []string{"a", "b", "c", "d"}
	show(letters)

	var s []string = make([]string, 5, 5)
	show(s)
}