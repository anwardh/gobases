package main

import "fmt"

func main() {

	a, b, c, d := 1, 0, 5, -3

	// Verificando se "a" é menor, maior ou igual a zero
	if a < 0 {
		fmt.Println("A é menor que 0") // A é maior que 0
	} else if a > 0 {
		fmt.Println("A é maior que 0")
	} else {
		fmt.Println("A é igual a 0")
	}

	// Verificando se "b" é menor, maior ou igual a zero
	if b < 0 {
		fmt.Println("B é menor que 0")
	} else if b > 0 {
		fmt.Println("B é maior que 0")
	} else {
		fmt.Println("B é igual a 0") // B é igual a 0
	}

	// Verificando se "c" é menor, maior ou igual a zero
	if c < 0 {
		fmt.Println("C é menor que 0")
	} else if c > 0 {
		fmt.Println("C é maior que 0") // C é maior que 0
	} else {
		fmt.Println("C é igual a 0")
	}

	// Verificando se "d" é menor, maior ou igual a zero
	if d < 0 {
		fmt.Println("D é menor que 0") // D é menor que 0
	} else if d > 0 {
		fmt.Println("D é maior que 0")
	} else {
		fmt.Println("D é igual a 0")
	}
}
