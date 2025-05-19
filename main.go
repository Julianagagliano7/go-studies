package main

import "fmt"

func main() {
	a := 10

	var pointer *int = &a //recupera o enderço de memória da variável a

	fmt.Println(pointer)
	fmt.Println(*pointer) //recupera o valor armazenado no endereço de memória

	*pointer = 50

	fmt.Println(*pointer)
	fmt.Println(pointer)
	fmt.Println(a)

	b := &a 
	*b = 20

	// 3 variáveis interligadas pelo apontamento ao mesmo endereço de memória 
	fmt.Println(*b)
	fmt.Println(*pointer)
	fmt.Println(a)
}
