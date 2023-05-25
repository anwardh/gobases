package main

import "fmt"

func main() {

	x := 15
	y := 10

	if x < y {
		fmt.Println("X é menor que Y")
	} else if x > y {
		fmt.Println("X é maior que Y")
	} else {
		fmt.Println("X é igual a Y") // X é igual a Y
	}

}
