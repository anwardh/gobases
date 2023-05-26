package main

import "fmt"

// Definição de uma Estrutura pessoa
type pessoa struct {
	nome      string // Campos e seus tipos de dados correspondentes
	idade     int
	profissao string
	peso      float64
}

func main() {

	// Instanciando uma nova estrutura - Forma 1 (Concisa)
	pessoa1 := pessoa{"Pessoa 1", 36, "Professor", 82.5}
	fmt.Println(pessoa1) // {Pessoa 1 36 Professor 82.5}

	// Instanciando uma nova estrutura - Forma 2 (campo e valor)
	pessoa2 := pessoa{
		nome:      "Pessoa 2",
		idade:     25,
		profissao: "Arquiteta",
		peso:      57.7,
	}

	fmt.Println(pessoa2) // {Pessoa 2 25 Arquiteta 57.7}

	// Acessando campos e valores específicos de uma estrutura
	fmt.Println(pessoa1.nome) // Pessoa 1

	// Alterando o valor de um campo específico de uma estrutura
	pessoa2.idade = 27
	fmt.Println(pessoa2) // {Pessoa 2 27 Arquiteta 57.7}

	// Definição e Instância de uma estrutura vazia
	var pessoa3 pessoa
	pessoa3.nome = "Pessoa 3"
	pessoa3.idade = 31
	pessoa3.profissao = "Médica"
	pessoa3.peso = 45.2

	fmt.Println(pessoa3) // {Pessoa 3 31 Zootecnista 45.2}
}
