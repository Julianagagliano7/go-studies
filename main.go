package main

import "fmt"

func main() {
	result := soma(20, 67)

	fmt.Println(result)
}

func soma (a, b int) int {
	return a + b
}