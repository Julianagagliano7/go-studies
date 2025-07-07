package main

import "fmt"

//range acessa o valor do elemento com mais facilidade

func main() {
	//Exemplo com slice
	// nums := []string{"pedro", "caio", "juliana"}

	// for key, value := range nums {
	// 	fmt.Println(key, value)
	// }

	//Exemplo com map
	users := map[string]string{
		"nome":      "João",
		"documento": "CPF",
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}
