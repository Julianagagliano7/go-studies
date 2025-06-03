package main

import "fmt"

//Map: estrutura chave-valor
//chave tem tipo, valor também

//Acesso rápido pois a chave é única

func main() {
	var pessoas = map[string]int{}
	pessoas["Ju"] = 21
	pessoas["Joao"] = 32
	pessoas["fabio"] = 25

	// fmt.Println(pessoas["Ju"])

	if idade, ok := pessoas["fabio"]; ok {
		fmt.Println("Idade do Fabio:", idade, ok)
	} else {
		fmt.Println("Fabio não encontrado")
	}

	//Removendo um elemento do map 
	delete(pessoas, "Ju")
	fmt.Println(pessoas)

	//Iterando um map 
	for nome, idade := range pessoas {
		fmt.Println("Nome:", nome, "Idade:", idade)
	}
}
