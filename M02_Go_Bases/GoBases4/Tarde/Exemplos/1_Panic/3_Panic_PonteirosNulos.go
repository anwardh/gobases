package main

import "fmt"

type Cachorro struct {
	nome string
}

func (c *Cachorro) latir() {
	fmt.Println(c.nome, "faz au au au")
}

func main() {

	// Instanciando c com a estrutura Dog para usar o método latir
	c := &Cachorro{"Katyta"}
	c = nil
	// Pânico gerado, em tempo de execução, pelo sistema
	c.latir()

}
