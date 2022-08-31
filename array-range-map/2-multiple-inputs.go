package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i])
	}

	fmt.Println(a)
}
