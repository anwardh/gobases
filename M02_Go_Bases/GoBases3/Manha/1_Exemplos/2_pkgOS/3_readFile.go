package main

import (
	"fmt"
	"os"
)

func main() {

	arquivo, erro := os.ReadFile("./2_readDir.go")

	if erro != nil {
		fmt.Println("Houve erro ao ler o arquivo", erro)
	}

	//fmt.Println(string(arquivo))
	fmt.Printf("%s", arquivo)

}
