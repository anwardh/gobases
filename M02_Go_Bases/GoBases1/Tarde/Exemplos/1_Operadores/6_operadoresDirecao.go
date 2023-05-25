package main

import "fmt"

func main() {

	var s string
	var ptr *string

	s = "ola"

	ptr = &s // p recebe o endereço da memória de s
	fmt.Println("O endereço de memória de s é:", &s)
	fmt.Println("O endereço de memória de ptr é:", &ptr)

	fmt.Println()

	fmt.Println("O valor de s é:", s)
	fmt.Println("O valor de ptr é:", *ptr) // p retorna o valor contido no endereço da memória de s

}
