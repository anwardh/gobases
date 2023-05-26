package main

import "fmt"

// Definição da função que retorna os quatro resultados das operações aritméticas
func operacores(valor1, valor2 float64) (float64, float64, float64, float64) {
	soma := valor1 + valor2
	subtracao := valor1 - valor2
	multiplicacao := valor1 * valor2

	var divisao float64
	if valor2 != 0 {
		divisao = valor1 / valor2
	}
	return soma, subtracao, multiplicacao, divisao

}

func main() {

	som, sub, mult, div := operacores(6, 2)

	// Execução da Função
	fmt.Println("Soma: \t\t", som)         // 8
	fmt.Println("Subtração: \t", sub)      // 4
	fmt.Println("Multiplicação: \t", mult) // 12
	fmt.Println("Divisão: \t", div)        // 3
}
