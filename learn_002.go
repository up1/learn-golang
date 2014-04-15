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

	buf := make([]byte, 1024)
	number, error := fileOpen.Write(buf)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("Success : %d\n",number)
}
