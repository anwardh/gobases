package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type produto struct {
	Nome      string
	Preco     float64
	Publicado bool
}

func main() {

	p := produto{
		Nome:      "MacBook Pro",
		Preco:     23.000,
		Publicado: true,
	}

	dadoEmJson, erro := json.Marshal(p)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(string(dadoEmJson))

}
