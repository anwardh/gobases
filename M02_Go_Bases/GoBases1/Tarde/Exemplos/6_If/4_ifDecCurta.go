package main

import "fmt"

func main() {

	if x := 1; x < 0 {
		fmt.Println("X é menor que Zero")
	} else if x > 0 {
		fmt.Println("X é maior que Zero")
	} else {
		fmt.Println("X é igual a Zero") // X é igual a Zero
	}

}
