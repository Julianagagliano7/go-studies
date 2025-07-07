package main

import (
	"fmt"
)

func main() {
	players := map[string]int{
		"lais": 26,
	}

	if value, ok := players["lais"]; ok {
		fmt.Println("pontos: ", value, ok)
	}

	if value, ok := players["juliana"]; !ok {
		fmt.Println("Não existe jogador com esse nome", value, ok)
	}
}
