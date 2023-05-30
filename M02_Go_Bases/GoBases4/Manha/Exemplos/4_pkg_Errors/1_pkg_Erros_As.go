package main

import (
	"errors"
	"fmt"
)

// Definição de uma estrutura para o erro
type meuErro struct {
	msg string
	x   string
}

// Definição do Método Error sobre a estrutura meuErro
func (m *meuErro) Error() string {
	return fmt.Sprintf("ocorreu um erro: %s, %s", m.msg, m.x)
}

func main() {

	// Instanciando a variável "e" e lhe atribuindo a estrutura meuErro com o seu retorno
	// Este erro simula o caso de ser o retorno de uma função que foi capturado, por exemplo
	// e queremos saber se esse erro corresponde a algum tipo especifico (meuErro)
	e := &meuErro{"novo error", "404"}

	// Criação da variável erro, do tipo meuErro
	var erro *meuErro

	temMeuErro := errors.As(e, &erro)

	fmt.Println(temMeuErro)
}
