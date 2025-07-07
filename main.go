package main

import "fmt"

func main() {
	nota := 65

	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	} else if nota >= 70 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
