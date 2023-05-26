package main

import "fmt"

// Declaração das constantes com as operações a serem realizadas
const (
	Soma          = "+"
	Subtracao     = "-"
	Multiplicacao = "*"
	Divisao       = "/"
)

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
	} else {
		return valor1 / valor2
	}
}

/* Definição da função que será encarregada de receber a operação a ser realizada e os valores aos quais a operação será aplicada
Para cada operação, vamos chamar uma função que recebe os valores e a função que iremos executar para aquele operador */
func operacaoAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Soma:
		return retornoOperacao(valores, operacaoSoma)
	case Subtracao:
		return retornoOperacao(valores, operacaoSubtracao)
	case Multiplicacao:
		return retornoOperacao(valores, operacaoMultiplicacao)
	case Divisao:
		return retornoOperacao(valores, operacaoDivisao)
	}
	return 0
}

// Definição da função que se encarregará de retornar as operações
// O primeiro parâmetro, recebe um slice de float64
// O segundo parâmetro é uma função que recebe 2 valores mas, que será executada pela função retornoOperacao neste ponto "resultado = operacao(resultado, valor)"
// operacao é um parâmetro (variável) do tipo func, que tem 2 valores do tipo float e retorna um float
func retornoOperacao(valores []float64, operacao func(valor1, valor2 float64) float64) float64 {
	var resultado float64           // Resultado irá acumular "a Soma"
	for i, valor := range valores { // O for fará um range em cima do array com os valores
		if i == 0 {
			resultado = valor // Se o índice for 0, resultado receberá esse único valor indicado no índice 0
		} else {
			resultado = operacao(resultado, valor) // Senão, resultado receberá a operação (segundo parâmetro), passando o valor de resultado e valor
		}
	}
	return resultado

}

func main() {

	// Execução da função
	fmt.Println("Soma: ", operacaoAritmetica(Soma, 10, 10, 10, 5, 9, 100))
	fmt.Println("Subtração: ", operacaoAritmetica(Subtracao, 10, 10, 10))
	fmt.Println("Multiplicação: ", operacaoAritmetica(Multiplicacao, 10, 10, 10))
	fmt.Println("Divisão: ", operacaoAritmetica(Divisao, 10, 10, 10))

}
