package main

import "fmt"

func main() {
	a := 10

	var pointer *int = &a //recupera o enderço de memória da variável a

	fmt.Println(pointer)
	fmt.Println(*pointer) //recupera o valor armazenado no endereço de memória
}
