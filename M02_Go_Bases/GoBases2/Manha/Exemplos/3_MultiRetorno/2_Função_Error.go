package main

import (
	"errors"
	"fmt"
)

// Definição de uma função que fará uma operação de divisão onde é validado se o divisor é zero,
// se for, retornará um erro, caso contrário, executará a divisão
func divisao(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("- O divisor não pode ser igual a zero")
	}
	return dividendo / divisor, nil
}

func main() {

	// Execução 1 da função
	//fmt.Println(divisao(2, 2)) // 0 - O divisor não pode ser igual a zero

	// Declaração/Inicialização das variáveis e Execução 2 da função
	resultado, erro := divisao(2, 2) // Divisor igual a 0
	if erro != nil {
		fmt.Println("Houve erro, resultado:", resultado, erro.Error()) // Houve erro, resultado: 0 - erro.Error() é a conversão do erro em String
	} else {
		fmt.Println("Não houve erro, resultado:", resultado)
	}

}
