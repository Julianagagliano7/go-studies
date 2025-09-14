package main

import (
	"fmt"
	"time"
)

func main() {
	//canal que recebe um inteiro
	ch := make(chan int, 3)

	//fica bloqueada até que outra goroutine leia o valor que ela escreveu, a não ser que o channel seja bufferizado
	go func() {
		ch <- 10
		ch <- 3
		ch <- 2
		ch <- 3 //não escreve esse valor até que um seja liberado (FIFO)
	}()

	time.Sleep(time.Second * 1)
	<-ch //esvazio o channel
	<-ch 
	valor := <-ch //leitura
	fmt.Println("valor do canal:", valor)

}
