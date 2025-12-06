package main

import "fmt"

//Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

func deferEx() {
	for i := 0; i < 5; i++ {
		if i < 3 {
			defer fmt.Println("i value", i)
		} else {
			fmt.Println(i)
		}
	}
}
func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred call") // runs at the end of main()

	fmt.Println("Middle")
	deferEx()
}
