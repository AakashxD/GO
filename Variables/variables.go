package main

import "fmt"

const value = 88

// not posible outside main function
//  hi:=89
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
