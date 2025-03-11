package main

import "fmt"
import "time"

func main() {
	for i := 1; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	// age := 19
	// for i := range 2 {
	// 	if age >= 18 {
	// 		fmt.Println("person is an adult")
	// 	} else {
	// 		fmt.Println("person is child")
	// 	}
	// 	fmt.Println("for range", i)
	// }
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("MOJ KRo MITRO")
	default:
		fmt.Println("WORK HARD YOU JUNKIEE")
	}
}
