package main

import "fmt"

func main() {
	variavel := 10

	fmt.Println(variavel)

	abc(&variavel)

	fmt.Println(variavel)
}

func abc(a *int) {
	*a = 200
}
