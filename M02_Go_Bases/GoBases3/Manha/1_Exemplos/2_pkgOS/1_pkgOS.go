package main

import (
	"log"
	"os"
)

func main() {

	erro := os.Setenv("Nome", "Gopher")

	// Checagem do erro
	if erro != nil {
		log.Println("Ocorreu um erro ao setar a variável de ambiente.", erro)
	}

	valor := os.Getenv("Nome")
	log.Println(valor)

	valor, existe := os.LookupEnv("Nomes")

	if existe {
		log.Println("A variável existe - valor: ", valor)
	} else {
		log.Println("A variável não existe")
	}

}
