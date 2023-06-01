package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type produto struct {
	Nome    string
	Preco   float64
	Postado bool
}

func main() {

	dadoJson := `{
		"Nome":"MacBook Pro",
		"Preco":23.575,
		"Postado":true
		}`

	// Instância de p como sendo uma estrutura produto (vazia)
	var p produto

	erro := json.Unmarshal([]byte(dadoJson), &p)

	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(p)
	fmt.Printf("%T", p)
}
