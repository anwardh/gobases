package main

import "fmt"

func olaMundo() {
	fmt.Println("Olá, Mundo!")
}

func main() {

	defer olaMundo()
	fmt.Println("Fim do programa")

}
