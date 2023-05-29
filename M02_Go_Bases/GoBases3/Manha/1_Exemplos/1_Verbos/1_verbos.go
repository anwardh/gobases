package main

import "fmt"

func main() {

	var nome string
	var idade int

	fmt.Println("Digite o seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Digite a sua idade: ")
	fmt.Scan(&idade)

	fmt.Printf("O nome digitado foi %s, com uma idade de %d anos.\n", nome, idade)

	fmt.Printf("O tipo da variável nome é %T\n", nome)
	fmt.Printf("O tipo da variável idade é %T", idade)
}
