package main

import (
	"fmt"
	"reflect"
)

/* Importação do pacote Reflect para fornecimento de funcionalidades acerca de obtenção de informaçõesdos objetos em tempo de execução.
Usado para descobrimos os tipos de elementos
*/

// Definição de uma Pessoa do tipo Estrutura com campos tipos de dados e rótulos para Banco de Dados
type Pessoa struct {
	PrimeiroNome string `bd:"primeiro_nome"`
	Sobrenome    string `bd:"sobrenome"`
	Telefone     string `bd:"telefone"`
	Endereco     string `bd:"endereco"`
}

func main() {

	// Forma de se obter o tipo 'reflect' da estrutura Pessoa
	pessoa := Pessoa{}          // pessoa recebe a estrutura Pessoa
	p := reflect.TypeOf(pessoa) // p recebe o "Tipo da estrutura (pessoa)"

	// Acesso e Visualização do nome da estrutura e o seu tipo
	fmt.Println("Type (Nome da Estrutura): ", p.Name()) // Type (Nome da Estrutura)  Pessoa
	fmt.Println("Kind (Tipo da Estrutura): ", p.Kind()) // Kind (Tipo da Estrutura):  Struct

	// Forma de se obter a quantidade de campos da estrutura - atrvés da Função NumField( )
	for i := 0; i < p.NumField(); i++ {
		// Obtendo o campo da estrutura passando o 'i' como parâmetro
		campo := p.Field(i)
		fmt.Println(i) // 0 1 2 3

		// Obtendo um valor dos rótulos dos campos
		tag := campo.Tag.Get("bd") // Rótulos da estrutura desenhada para o Banco de Dados
		fmt.Println("\t", tag)     // primeiro_nome sobrenome telefone endereco
	}

}
