package main

import (
	"fmt"
	"os"
)

func main() {

	fileOpen, error := os.Create("test.txt")
	if error != nil {
		fmt.Println(error)
	}
	defer func() {
		fmt.Println("defer")
		if error := fileOpen.Close(); error != nil {
			fmt.Println(error)
		}
	}()
}
