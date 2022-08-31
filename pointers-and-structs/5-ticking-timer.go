package main

import "fmt"

/* Define a Timer struct
with two fields: id and value */
type Timer struct {
	id    string
	value int
}

/* define the tick() method, which
should increment the value by one
and output its current value. */
func (t *Timer) tick() {
	for i := 0; i < t.value; i++ {
		fmt.Println(i + 1)
	}
}

func main() {
	var x int
	fmt.Scanln(&x)
	t := Timer{"timer1", x}
	t.tick()

}
