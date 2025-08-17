package main

import "fmt"

// & -> imprime o enderço
// * -> imprime o conteúdo

//Com ponteiros consigo modificar o valor original da variável

type Pessoa struct {
	Nome string
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "Juliana"}
	var p2 Pessoa = Pessoa{Nome: "Esmael"}
	fmt.Println(p1)

	var p3 *Pessoa = &p1
	p3.Nome = "Joana"

	// fmt.Println(&p1.Nome)
	// fmt.Println(&p3.Nome)
	
	fmt.Println(p1)
	fmt.Println(p2)
}
