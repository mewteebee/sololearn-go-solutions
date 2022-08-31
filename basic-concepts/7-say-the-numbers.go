package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	strNums := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	aStr := strNums[a]
	bStr := strNums[b]
	cStr := strNums[c]

	fmt.Println(aStr)
	fmt.Println(bStr)
	fmt.Println(cStr)
}
