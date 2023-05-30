package main

import (
	"fmt"
	"log"
)

func main() {

	defer func() {
		fmt.Println("Mensagem após o Log Fatal")
	}()

	log.Fatal("Log Fatal - Finalização do Programa")

	panic("Panic após o Log Fatal")

}
