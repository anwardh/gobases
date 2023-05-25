package calc2

import "fmt"

func Multiplicacao(a, b int) int {
	return a * b
}

func Divisao(a, b int) int {

	if b == 0 {
		fmt.Println("Não existe divisão por Zero!")
	}

	return a / b
}
