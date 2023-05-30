package main

import (
	"errors"
	"fmt"
	"log"
)

// Definição da função que receberá uma string e retornará dois valores (string e erro)
func DizerOlar(nome string) (string, error) {
	// Se o nome for vazio, a função retornará uma mensagem de erro
	if nome == "" {
		return "", errors.New("\nO nome está vazio")
	}
	// Se o valor do nome não for vazio, retornará uma mensagem de boas-vindas
	return fmt.Sprintf("Olá, %s", nome), nil
}

func main() {

	// Chamada da função passando um valor para o parâmetro nome
	nome, erro := DizerOlar("Thamyris")

	// Checagem do erro
	if erro != nil {
		log.Println(erro) // Inmpressão da mensagem de erro e um log
	}

	fmt.Println(nome) // Impressão do nome
}
