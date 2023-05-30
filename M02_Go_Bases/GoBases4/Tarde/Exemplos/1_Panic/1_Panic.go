package main

// Definição de um programa em que será realizada a leitra de um arquivo semArquivo.txt.
import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Iniciando...")

	_, erro := os.Open("semArquivo.txt")

	// Se o erro for diferende de nulo, será gerado um panic
	if erro != nil {
		panic(erro) // panic: open semArquivo.txt: O sistema não pode encontrar o arquivo especificado.
		//fmt.Println(erro)
		// exit status 2 -> A execução foi interrompida
	}

	fmt.Println("Fim do Programa")

}
