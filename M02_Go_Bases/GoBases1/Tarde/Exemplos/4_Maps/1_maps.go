package main

import "fmt"

func main() {

	estudantes := map[string]int{

		"João":  21,
		"Maria": 19,
	}

	fmt.Println(estudantes) // map[João:21 Maria:19]

	// Adicionando elementos
	estudantes["Pedro"] = 22
	estudantes["Eduarda"] = 18

	fmt.Println(estudantes) // map[Eduarda:18 João:21 Maria:19 Pedro:22]

	// Atualizando valores
	estudantes["João"] = 23

	fmt.Println(estudantes) // map[Eduarda:18 João:23 Maria:19 Pedro:22]

	// Deletando elementos
	delete(estudantes, "Pedro")

	fmt.Println(estudantes) // map[Eduarda:18 João:23 Maria:19]

	// // Imprimindo o Comprimento do Map
	fmt.Println(len(estudantes)) // 2
	fmt.Println()

	// Recorrendo elementos (Chave e Valor) de um Map
	for chave, valor := range estudantes {

		fmt.Println("Chave:", chave, " - ", "Valor:", valor)

	}

	/*

		Chave: João  -  Valor: 23
		Chave: Maria  -  Valor: 19
		Chave: Eduarda  -  Valor: 18

	*/

}
