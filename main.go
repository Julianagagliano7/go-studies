package main

import "fmt"

//slice possuem tamanho flexível (dinâmico)
//armazenam dados do mesmo tipo

//inserção: uso a função append (aumenta o slice por baixo dos panos), insere sempre no final

func main() {
	var gavetas []string

	gavetas = append(gavetas, "copos", "panos", "colheres") //inserção
	fmt.Println(gavetas[2])                                 //leitura

	fmt.Println(len(gavetas))

	// Divisão de slice
	// slice[x:x-1] -> último x é sempre o len do slice - 1

	fmt.Println(gavetas[0:2])


	//Remoção de itens do slice 

	gavetas = gavetas[:2] 
	fmt.Println(gavetas) 
}
