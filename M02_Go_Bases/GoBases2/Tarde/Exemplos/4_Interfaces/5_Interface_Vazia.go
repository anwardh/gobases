package main

import "fmt"

// Definição de uma Estrura com um campo Data sendo do tipo Interface (vazia)
type ListaHeterogenea struct {
	Data []interface{}
}

func main() {

	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "Olá")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data) // [1 Olá true]

}
