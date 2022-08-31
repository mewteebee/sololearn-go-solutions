package main

import "fmt"

func route(a []string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(a[i] + "->")
	}
}

func main() {
	var n int
	fmt.Scanln(&n) // size of slice

	a := make([]string, n)
	for j := 0; j < n; j++ { // populate slice
		fmt.Scanln(&a[j])
	}
	route(a, n)
}
