package main

import "fmt"

//Declaração da função isEven com um defer que permite que, mesmo tendo um panic, ela será executada
func isEven(numero int) {
	/* A defer pode ser usada em casos de panic e, a função a seguir é uma forma de se trabalhar com ela ou,
	poderíamos declará-la como escopo de uma outra função e chamar esta no local em que a defer deverá ser usada.*/

	defer func() {
		// A função recover está sendo alocada na variável erro
		erro := recover()
		// Se houver esse erro, será printado na tela
		// Dessa forma, o programa não mais será interrompido de vez, porque conseguios nos recuperar
		// nos casos de erro
		if erro != nil {
			fmt.Println(erro)
		}
	}()

	if (numero % 2) != 0 {
		// o panic é gerado e seu valor é recuperado pela recover que irá imprimí-lo e continuar com o programa
		//panic(nil)
		panic("É Ímpar")

	}

	fmt.Println(numero, "é Par")
}

func main() {
	isEven(4)

	fmt.Println("Fim do Programa")
}
