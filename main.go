package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Pessoa é uma struct que recebe o método Apresentar
// Método apresentar fica exclusivamente relacionado à essa struct e pode modificar seus valores e comportamentos

//func (p Pessoa) -> recebem uma cópia da struct (não a struct original), com isso, a struct original não é modificada
//func (p *Pessoa) -> recebem um ponteiro da struct original, logo, ela é modificada

//obs: remova o asterisco e rode para ver a diferença entre os cenários

func (p *Pessoa) Apresentar() {
	p.Nome = "Roberto" //modifica a struct original
	fmt.Printf("Olá, meu nome é %s e tenho e tenho %d anos\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Juliana", Idade: 21}
	p1.Apresentar()
	fmt.Println(p1.Nome)
}
