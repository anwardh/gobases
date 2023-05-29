package main

import (
	"fmt"
	"os"
)

func main() {

	conteudo := []byte("Funcionou!")

	erro := os.WriteFile("./ola.txt", conteudo, 0644)

	if erro != nil {
		fmt.Println("Erro ao escrever o arquivo", erro)
	}

}
