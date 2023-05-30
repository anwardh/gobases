package main

import (
	"errors"
	"fmt"
)

func main() {
	statusCode := 400
	// Checagem do erro
	if statusCode > 399 {
		// Uso da função New( ) para retorno do erro
		fmt.Println(errors.New("A requisição falhou."), statusCode)
		return
	}
	fmt.Println("O programa foi encerrado corretamente.")
}
