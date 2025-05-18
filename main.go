package main

import "fmt"

func main() {
	var floatNumber float32 = 1.1 //32 bits - apenas operações matemáticas do mesmo tipo de float são permitidas

	var pi float64 = 3.14159
	var radius float64 = 5.0
	var area float64 = pi * radius * radius

	fmt.Println("Area of the circle:", area)
	fmt.Println("Float number:", floatNumber)
}
