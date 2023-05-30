package main

import (
	"errors"
	"fmt"
)

// Definição da variável erro1 que retorna um erro
var erro1 = errors.New("erro número 1")

// Na execução, var2 dará em "false" porque, na função "x", o retorno é para o "erro1"
var erro2 = errors.New("erro número 2")

// Definição da função "x" que retorna um erro
func x() error {
	return fmt.Errorf("informação extra do erro: %w", erro1)
}

func main() {

	// Instanciando a variável "e"
	e := x()

	// Comparando se "e", recebido da função é igual a variável "erro2", mas não é
	// por conta do retorno da função "x", que é "erro1"
	coincidence := errors.Is(e, erro1) //
	fmt.Println(coincidence)

}
