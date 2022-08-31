package main

import "fmt"

type Cart struct {
	prices []float32
}

func (x Cart) show() {
	var sum float32 = 0.0
	//calculate the sum of all prices in the Cart
	for _, v := range x.prices {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	c := Cart{}
	var n int
	fmt.Scanln(&n)
	var x float32
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		c.prices = append(c.prices, x)
	}

	// take n inputs and add them to the Cart

	//call the show() method of the Cart
	c.show()
}
