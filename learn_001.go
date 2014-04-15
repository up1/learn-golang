package main

import (
	"fmt"
)

func main() {
	i := 1
	if i == 1 {
		fmt.Println("test")
	}

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

}
