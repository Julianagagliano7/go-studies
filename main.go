package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() { //ponteiro altera o valor de carro.Name na main()
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {
	carro := Carro{
		Name: "Citroen C3",
	}

	carro.andou()
	fmt.Println(carro.Name)
}
