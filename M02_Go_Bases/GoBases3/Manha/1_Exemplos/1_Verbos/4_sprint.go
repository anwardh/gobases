package main

import "fmt"

func main() {

	nome, idade := "Lucas Ganda", 23

	res := fmt.Sprint(nome, " tem ", idade, " anos de idade")
	fmt.Println(res)
	fmt.Printf("O Tipo de \"res\" é: %T\n", res)

	res = fmt.Sprintf("%s tem %d anos de idade.", nome, idade)
	fmt.Println(res)
}
