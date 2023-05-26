package main

import "fmt"

// Definição da Função variádica
func soma(valores ...float64) float64 {
	// Declaração da variável que irá receber o somatório dos valores inseridos
	var resultado float64
	// Estrutura que percorrerá todos os 'valores' inseridos
	for _, valor := range valores {
		resultado += valor // resultado recebe e acumula o valor da soma de cada um dos valores inseridos
	}
	return resultado // Retorno da função
}

func main() {

	// Chamada da função e sua impressão na tela
	fmt.Println(soma(2, 3, 2, 1, 2, 3, 4, 5, 6)) // 28

}
