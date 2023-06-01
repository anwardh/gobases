package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func raizHandler(context *gin.Context) {

	//O corpo, cabeçalho e método estão contidos no contexto gin.
	body := context.Request.Body
	header := context.Request.Header
	method := context.Request.Method

	fmt.Println("Eu recebi algo!")           // Imprime no console...
	fmt.Printf("\tMétodo: %s\n", method)     // ... qual o método utilizado e ...
	fmt.Printf("\tConteúdo do cabeçalho:\n") // ... o conteúdo do cabeçalho

	// Imprimindo todos os metadados presentes no Cabeçalho
	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	// Mensagem contida no Body
	fmt.Printf("\tO body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que ler depois\n", body)
	fmt.Println("¡Yay!")
	// Retorna a string seguinte quando a conexão é bem sucedida
	context.String(http.StatusOK, "Eu recebi!") //Respondemos ao cliente com 200 OK e uma mensagem.
}

func main() {

	r := gin.Default()
	r.GET("/", raizHandler)
	r.Run()
}
