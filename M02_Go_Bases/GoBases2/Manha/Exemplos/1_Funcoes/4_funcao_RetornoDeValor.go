package main

import "fmt"

// Declaração de uma função com dois valores do tipo float64 e que retorna um valor, também, do tipo float64

func soma(valor1, valor2 float64) float64 {
	return valor1 + valor2 // Definição do retorno da função
}

func main() {

	// Instanciando a variável s com a chamada da função e seus devidos argumentos
	s := soma(4, 5)

	// Impressão da variável s na tela com o resultado retornado pela função Soma
	fmt.Println(s) // 9

}
