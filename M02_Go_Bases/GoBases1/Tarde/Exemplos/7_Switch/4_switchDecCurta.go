package main

import "fmt"

func main() {

	switch dia := "domingo"; dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Printf("%s é um dia de final de semana\n", dia) // domingo é um dia de final de semana

	}

}
