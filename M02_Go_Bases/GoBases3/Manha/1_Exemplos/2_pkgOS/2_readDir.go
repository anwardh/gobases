package main

import (
	"fmt"
	"os"
)

func main() {

	diretorio, erro := os.ReadDir("../2_pkgOS")

	// Checagem do erro

	if erro != nil {
		fmt.Println("Erro ao ler o diretório", erro)
	}

	//fmt.Println(diretorio)

	fmt.Printf("%T\n", diretorio)

	for _, arquivos := range diretorio {
		fmt.Println(arquivos.Name())

		fmt.Printf("%s é uma pasta? %t\n", arquivos.Name(), arquivos.Type().IsDir())
	}

}
