package main

import "fmt"

func main() {

	frutas := []string{"Uva", "Maçã", "Maracujá", "Melão", "Banana"}
	for indice, valor := range frutas {
		fmt.Println("Índice:", indice, " - ", "Fruta:", valor)
	}

}

/*

	Índice: 0  -  Fruta: Uva
	Índice: 1  -  Fruta: Maçã
	Índice: 2  -  Fruta: Maracujá
	Índice: 3  -  Fruta: Melão
	Índice: 4  -  Fruta: Banana

*/
