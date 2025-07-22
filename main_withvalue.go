package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		"testKey", "testValue",
	)

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	fmt.Println(ctx.Value("testKey"))

	//alterando o valor da do contexto criando outro contexto e passando o existente 
	ctx2 := context.WithValue(ctx, "testKey", "juliana")
	fmt.Println(ctx2.Value("testKey"))
}
