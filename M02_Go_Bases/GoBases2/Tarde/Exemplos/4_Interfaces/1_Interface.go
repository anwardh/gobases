package main

import "fmt"

type Animal interface {
	Som() string
	Movimento() string
}

type Cachorro struct {
	Nome string
}

// O método está implementando a esrutura Cachorro
func (c Cachorro) Som() string {
	return "Au au au"
}

func (c Cachorro) Movimento() string {
	return "Correr"
}

func main() {

	var cachorro Animal
	cachorro = Cachorro{Nome: "Chico"}
	fmt.Println(cachorro)             // {}
	fmt.Println(cachorro.Som())       // Au au au
	fmt.Println(cachorro.Movimento()) // Correr
}
