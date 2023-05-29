package main

import "fmt"

func main() {

	nome, idade := "Jainy", 18

	fmt.Print(nome, " tem ", idade, " anos.\n")
	fmt.Printf("%s tem %d anos.\n", nome, idade)
	fmt.Println(nome, "tem", idade, "anos.")

}
