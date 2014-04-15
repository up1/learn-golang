package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	letters = append(letters, "e", "f")
	fmt.Println(letters)


	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}
