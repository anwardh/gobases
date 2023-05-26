package main

import "fmt"

// Definição da Função com seus valores de retorno nomeados com suas respectivas variáveis
func operacores(valor1, valor2 float64) (soma, subtracao, multiplicacao, divisao float64) {
	soma = valor1 + valor2
	subtracao = valor1 - valor2
	multiplicacao = valor1 * valor2

	if valor2 != 0 {
		divisao = valor1 / valor2
	}
	return
}

func main() {

	// Execução da Função - Modo 1
	//fmt.Println(operacores(2, 2)) // 4 0 4 1

	// Execução da Função - Modo 2
	soma, subtracao, multiplicacao, divisao := operacores(10, 2)
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtracao)
	fmt.Println("Multiplicação: ", multiplicacao)
	fmt.Println("Divisão: ", divisao)

}
