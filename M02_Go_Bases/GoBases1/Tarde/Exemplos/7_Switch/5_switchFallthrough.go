package main

import (
	"fmt"
	"time"
)

func main() {

	hoje := time.Now()
	var t int = hoje.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Limpar a casa")
	case 25, 26, 27:
		fmt.Println("Comprar comida")
		fallthrough
	case 31:
		fmt.Println("Hoje tem festa")
		fallthrough
	case 29:
		fmt.Println("Hoje tem animação")
		fallthrough
	default:
		fmt.Println("Não há informações disponíveis para esse dia")

	}

}
