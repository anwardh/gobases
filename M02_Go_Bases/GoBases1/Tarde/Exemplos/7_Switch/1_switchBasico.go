package main

import "fmt"

func main() {

	dias := 1

	switch dias {
	case 0:
		fmt.Println("Segunda-feira")
	case 1:
		fmt.Println("Terça-feira") // Terça-feira
	case 2:
		fmt.Println("Quarta-feira")
	case 3:
		fmt.Println("Quinta-feira")
	case 4:
		fmt.Println("Sexta-feira")
	case 5:
		fmt.Println("Sábado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Desconhecido")
	}
}
