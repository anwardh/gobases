package main

import "fmt"

func ola() {
	fmt.Println("Olá, eu sou uma Go Routine")
}

func main() {

	go ola()
	fmt.Println("Estou dentro da Func Main")
}
