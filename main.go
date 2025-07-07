package main

import (
	"fmt"
	"time"
)

// switch condicao {
// 	case valor1:
//	default:
//}

func main() {
	fmt.Println("Quando é sábado?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("é hoje")
	case today + 1:
		fmt.Println("é amanhã")
	case today + 2:
		fmt.Printf("é em dois dias")
	default:
		fmt.Println("Tá longe ainda...")
	}
}
