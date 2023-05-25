package main

import "fmt"

func main() {

	dia := "domingo"
	switch dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Println("É um dia de final de semana:", dia) // domingo é um dia de final de semana
	}

}
