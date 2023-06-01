package main

import (
	"fmt"
	"net/http"
)

// Esta função executará a ação de printar na tela quando acessarmos o caminho http://localhost:8080/ola
func dizerOla(w http.ResponseWriter, req *http.Request) {
	// Mensagem formatada a ser apresentada na tela
	fmt.Fprintf(w, "Olá, está funcionando :)\n")
}

func main() {

	http.HandleFunc("/ola", dizerOla)
	http.ListenAndServe(":8080", nil)
}
