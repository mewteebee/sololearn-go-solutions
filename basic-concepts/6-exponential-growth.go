package main

import "fmt"

func main() {
	var years int
	var res int = 7
	fmt.Scanln(&years)

	for i := 0; i < years; i++ {
		res = res * 2

	}

	fmt.Println(res)
}
