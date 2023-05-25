package main

import "fmt"

func main() {

	idade := 16

	switch {
	case idade >= 110:
		fmt.Println("É importal?")
	case idade >= 18:
		fmt.Println("É maior de idade") // É maior de idade
	default:
		fmt.Println("É menor de idade")
	}
}
