package main

import "fmt"

func main() {

	var i interface{} = "hello"

	// Verificação se é string
	s := i.(string)
	fmt.Println(s) // hello

	// Verificação se é string e "ok"
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	// Verifica se é float
	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic : o validador (ok) não foi adicionado
	fmt.Println(f)
}
