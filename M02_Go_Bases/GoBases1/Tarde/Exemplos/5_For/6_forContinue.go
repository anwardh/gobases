package main

import "fmt"

func main() {

	for i := 0; i < 20; i++ {

		if i%2 == 0 { // Verifica se o módulo da divisão de i por 2 é igual a 0
			continue // Se a condição for falsa, o 'continue' passará para a próxima iteração
		}
		fmt.Println(i) // 1 3 5 7 9 11 13 15 17 19

	}
}
