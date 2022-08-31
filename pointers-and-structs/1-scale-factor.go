package main

import "fmt"

func scale(x *float32, y float32) {
	*x = *x * y
}

func main() {
	var num float32
	var factor float32

	fmt.Scanln(&num)
	fmt.Scanln(&factor)

	scale(&num, factor)
	fmt.Println(num)
}
