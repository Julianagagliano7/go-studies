package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello string = "Hello, World!\n"
	var question string = "How are you?"

	var meet = hello + question

	fmt.Println(strings.ToUpper(meet))
	fmt.Println(strings.Contains(meet, "World"))
	fmt.Println(strings.Contains(meet, "world"))
}
