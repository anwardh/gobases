package main

import "fmt"

func main() {

	var a [2]string // Declaração de um Array de tamanho 2 e do tipo string
	a[0] = "Hello"  // Atribuição de valor ao elemento "0" do array
	a[1] = "World"

	fmt.Println(a[0], a[1]) // Hello World - Impressão isolada dos valores dos elementos do array
	fmt.Println(a)          // [Hello World] - Impressão do array

}
