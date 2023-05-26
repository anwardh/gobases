package main

import "fmt"

// Definição das variáveis referentes às Operações Aritéticas
const (
	Soma          = "+"
	Subtracao     = "-"
	Multiplicacao = "*"
	Divisao       = "/"
)

// Definição da Função que vai lidar com o tipo de operação que indicarmos
func operacaoAritmetica(operador string, valor1, valor2 float64) float64 {

	switch operador {
	case Soma:
		return valor1 + valor2
	case Subtracao:
		return valor1 - valor2
	case Multiplicacao:
		return valor1 * valor2
	case Divisao:
		if valor2 != 0 {
			return valor1 / valor2
		} else {
			fmt.Println("Não existe divisão por 0!")
		}
	}
	return 0 // Retorno da Função
}

func main() {

	// Chamada/Execução da função com cada uma das operações
	fmt.Println(operacaoAritmetica(Soma, 6, 2))          // 8
	fmt.Println(operacaoAritmetica(Subtracao, 6, 2))     // 4
	fmt.Println(operacaoAritmetica(Multiplicacao, 6, 2)) // 12
	fmt.Println(operacaoAritmetica(Divisao, 6, 0))       // 3
}
