package main

import "fmt"

func main() {
	var idade int = 30   //ocupa tamanho da memória do sistema, padrão 64 bits
	var idade1 int32 = 2 //ocupa 32 bits
	var idade2 int8 = 1 //ocupa 8 bits

	fmt.Println(idade)
	fmt.Println(idade1)
	fmt.Println(idade2)
}

