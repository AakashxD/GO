package main

import (
	"fmt"
	"math"
)

// that the type comes after the variable name In arguments.

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

//When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

func sqrt(x float64) string {
	//Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, 4); x > 400 {
		return fmt.Sprint(v)
	}

	if x < 0 {
		return sqrt(-x) + "i"
	}
	// .Sprint to convert into String
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println("Main func")
	var a = add(4, 5)
	fmt.Println(a)
	c, d := swap("aakash", "singh")
	fmt.Println(c, d, sqrt(-4))

}
