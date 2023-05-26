package main

import "fmt"

// Definição de uma Esrutuda Preferências
type preferencias struct {
	comidas  string
	filme    string
	series   string
	animes   string
	esportes string
}

// Definição de uma Estrutura Pessoa
type pessoa struct {
	nome      string // Campos e seus tipos de dados correspondentes
	idade     int
	profissao string
	peso      float64
	gostos    preferencias
}

func main() {

	// Instanciando uma nova estrutura
	pessoa1 := pessoa{
		nome:      "Pessoa 1",
		idade:     36,
		profissao: "Professor",
		peso:      82.5,
		// Adicionando os campos e valores da estrutura preferencias dentro da estrutura pessoa
		gostos: preferencias{
			comidas:  "Feijoada",
			filme:    "Curtindo a Vida Adoidado",
			series:   "Breaking Bad",
			animes:   "Yu Yu Hakusho",
			esportes: "Trilhas",
		},
	}

	fmt.Println(pessoa1) // {Pessoa 1 36 Professor 83.5 {Feijoada Curtindo a Vida Adoidado Breaking Bad Yu Yu Hakusho Trilhas}}

	// Alterando os valores de campos específicos
	pessoa1 = pessoa{
		nome:      "Pessoa 1",
		idade:     34,
		profissao: "Arquiteta",
		peso:      52.2,
		gostos: preferencias{
			comidas: "Macarronada",
			filme:   "Clube dos Cinco",
			//series:   "Vikings",
			animes:   "Sailor Moon",
			esportes: "Natação",
		},
	}

	fmt.Println(pessoa1) // {Pessoa 1 37 Arquiteta 52.2 {Macarronada Clube dos Cinco  Sailor Moon Natação}}

	// Alterando os valores de campos específicos
	pessoa1.gostos.comidas = "Farofa de ovo"
	fmt.Println(pessoa1) // {Pessoa 1 37 Arquiteta 52.2 {Farofa de ovo Clube dos Cinco  Sailor Moon Natação}}

}
