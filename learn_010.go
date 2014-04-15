package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters[1])
	fmt.Println(letters[:])
	fmt.Println(letters[:4])
	fmt.Println(letters[1:2])
	fmt.Println(letters[1:4])
	fmt.Println(letters[1:])

}
