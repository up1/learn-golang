package main

import "fmt"

func main() {
	teams := map[string]int {
		"MAN U": 20,
		"MANCI": 30,
	}

	fmt.Println(teams)

	fmt.Println(teams["MAN U"])

	var score int
	var ok bool
	score, ok = teams["MANU"]
	fmt.Println(score)
	fmt.Println(ok)
}
