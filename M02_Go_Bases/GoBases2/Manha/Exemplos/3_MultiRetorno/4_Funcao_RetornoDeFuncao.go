package main

import "fmt"

// Definição das funções para cada operação
func operacaoSoma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func operacaoSubtracao(valor1, valor2 float64) float64 {
	return valor1 - valor2

}
func operacaoMultiplicacao(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func operacaoDivisao(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2

}

// Definição da uma função que se encarrega de lidar com as funções que realização as operações

func operacaoAritmetica(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Soma":
		return operacaoSoma
	case "Subtracao":
		return operacaoSubtracao
	case "Multiplicacao":
		return operacaoMultiplicacao
	case "Divisao":
		return operacaoDivisao
	}
	return nil

}

func main() {

	/* Declaração das variáveis que recebem a operação aritmética a ser realizada, bem como, os valores	que serão utilizados em cada uma delas */
	// operacao está recebendo a operacaoAritmetica que irá executar a função operacaoSoma
	// resultado está recebendo a operacaoSoma (indicada anteriormente) usando os argumentos 2 e 7
	//operacao := operacaoAritmetica("Soma")
	//resultado := operacao(2, 5)
	fmt.Println(operacaoAritmetica("Soma")(2, 5))
	fmt.Println(operacaoAritmetica("Subtracao")(2, 5))
	fmt.Println(operacaoAritmetica("Multiplicacao")(2, 5))
	fmt.Println(operacaoAritmetica("Divisao")(6, 2))

}
