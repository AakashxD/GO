package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(30), math.Sqrt(4))

	fmt.Println(math.Pi) // math.pi will throw error ->In Go, a name is exported if it begins with a capital letter.
}
