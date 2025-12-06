package main

import (
	"fmt"
	"math"
)

//Constants cannot be declared using the := syntax.

const value = 88

/*
*
Go provides signed (int, int8–int64) and unsigned integers (uint, uint8–uint64) for whole numbers of different sizes.

	byte is an alias for uint8 and used for raw data, while rune is an alias for int32 and represents Unicode characters.
	float32 and float64 handle decimal values, with float64 being the default for precision.

complex64 and complex128 store complex numbers, and uintptr holds memory addresses for low-level operations.

	not posible outside main function
	 hi:=89

/

/**
Variables declared without an explicit initial value are given their zero value.

	The zero value is:
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
*/

func typeCasting() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
func main() {
	var name = "Aakash"
	surname := "singh"
	fmt.Println(name, surname)

	const (
		PORT = 9999
		HOST = "http://localhost:9999"
	)
	a := 10.0
	b := 3.0
	fmt.Println(PORT, HOST, value, a/b)
}
