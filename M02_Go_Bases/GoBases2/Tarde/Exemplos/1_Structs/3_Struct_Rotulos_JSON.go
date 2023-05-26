package main

import (
	"encoding/json"
	"fmt"
)

// Definição de uma Pessoa do tipo Estrutura com campos, tipos de dados e rótulos para formato JSON
type Pessoa struct {
	PrimeiroNome string `json:"primeiro_nome"`
	Sobrenome    string `json:"sobrenome"`
	Telefone     string `json:"telefone"`
	Endereco     string `json:"endereco"`
}

func main() {

	// Declaração/Inicialização de uma estrutura p do tipo pessoa

	p := Pessoa{
		PrimeiroNome: "Pessoa1",
		Sobrenome:    "Sobrenome1",
		Telefone:     "88888888888",
		Endereco:     "Alameda dos Anjos",
	}

	/* formatoJSON recebe a conversão dos dados em JSON por meio da função .Marshal (que retorna os bytes da representação dodificada em JSON)
	   "erro" retorna uma possível mensagem de erro caso haja falha na conversão*/
	formatoJSON, erro := json.Marshal(p)

	// CHECAGEM DO ERRO
	if erro != nil {
		fmt.Println("Não foi possível gerar o JSON. erro:", erro)
	}

	// Conversão em String
	fmt.Println(string(formatoJSON)) // {"primeiro nome":"Pessoa1","sobrenome":"Sobrenome1","telefone":"88888888888","endereco":"Alameda dos Anjos"}
	fmt.Println(erro)                // <nil>
}
