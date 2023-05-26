package main

import "fmt"

// Declaração da Função com seu parâmetro
func verificarVariavel(numero int) {

	// Definição do Escopo
	if numero < 0 {
		fmt.Println("O número é menor que 0")
	} else if numero > 0 {
		fmt.Println("O número é maior que 0")
	} else {
		fmt.Println("O número é igual a 0")
	}

}

func main() {

	// Declaração/Inicialização das variáveis dentro da Função Main
	a, b, c, d := 1, 0, 5, -3

	// Execução da Função verificarVariável para cada uma das variáveis
	verificarVariavel(a) // O número ( a ) é maior que 0
	verificarVariavel(b) // O número ( b ) é igual a 0
	verificarVariavel(c) // O número ( c ) é maior que 0
	verificarVariavel(d) // O número ( d ) é menor que 0

}
