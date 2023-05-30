package main

import "fmt"

func main() {
	//aplicamos “defer” à invocação de uma função anônima
	defer func() {
		//fmt.Println("Esta função será executada mesmo gerando um panic")
		fmt.Println("Esta função será executada mesmo gerando um panic")

	}()

	fmt.Println("Mensagem entre o Defer e o Panic")
	//criamos um panic com a sua mensagem de retorno
	// panic("Panic criado!!!")

}
