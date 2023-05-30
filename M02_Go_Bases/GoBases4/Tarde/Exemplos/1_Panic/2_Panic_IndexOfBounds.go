package main

import "fmt"

func main() {
	animais := []string{
		"cachorro",
		"bode",
		"urubu",
	}
	// Neste caso, o erro está por conta do acesso a um índice fora fora do comprimento do slice.
	//fmt.Println("O animal que voa é o:", animais[len(animais)]) // panic: runtime error: index out of range [3] with length 3
	fmt.Println("O animal que voa é o:", animais[len(animais)-1])

	fmt.Println("Fim do programa")
}
