package main

import "fmt"

func main() {
	var v int = 19

	var p *int

	p = &v

	fmt.Println("O endereço de memória de V é:", &v)
	fmt.Printf("O ponteiro P refere-se ao endereço de memória: %v\n", p)
	fmt.Println("Ao Desreferenciar o valor de P, temos: ", *p)

}
