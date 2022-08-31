package main

import "fmt"

func main() {
	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	var x int
	fmt.Scanln(&x)
	if x > 5 || x < 0 {
		fmt.Println("Invalid choice")
	} else {
		fmt.Println(menu[x])
	}

}
