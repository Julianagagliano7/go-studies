package main

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {

	req := Request{
		Name: "Juliana",
		Age:  20,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Erro ao fazer marshal: ", err)
	}

	fmt.Println("Resutado - Marshal: ", payload)

	jsonString := `{"nome":"João","idade":30}`
	var p Pessoa

	err = json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Erro ao fazer unmarshal: ", err)
		return
	}
	fmt.Println("Resultado Unmarshal: ", p)

}
