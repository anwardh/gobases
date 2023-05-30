package main

import (
	"errors"
	"fmt"
)

// Criação da estrutura erroDois
type erroDois struct{}

// Definição do método Error para a estrutura erroDois
func (e erroDois) Error() string {
	return "erro 2 aconteceu"
}

func main() {

	// Instanciando e2 com a estrutura erroDois
	e2 := erroDois{}

	// Instanciando e1 retornando um erro, tendo como parâmetro e2
	e1 := fmt.Errorf("e2: %w", e2)

	// Quando a função executar o e1, ela irá verificar se o e2 está contido
	fmt.Println(errors.Unwrap(e1))
	fmt.Println(errors.Unwrap(e2))

}
