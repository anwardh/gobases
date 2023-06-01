package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definição da pseudo 'base de dados' de empregados
var empregados = map[string]string{
	"1": "Empregado A",
	"2": "Empregado B",
	"3": "Empregado C",
	"4": "Empregado D",
}

// Handler responsável por responder no endpoint "/" (raiz)
func paginaRaiz(ctxt *gin.Context) {
	ctxt.String(http.StatusOK, "Página Principal")
}

// Handler que verificará se o id passado pelo cliente existe na base de dados
func acharEmpregado(ctxt *gin.Context) {
	empregado, ok := empregados[ctxt.Param("id")]
	if ok {
		ctxt.String(http.StatusOK, "Informação do Empregado %s, nome: %s", ctxt.Param("id"), empregado)
	} else {
		ctxt.String(http.StatusNotFound, "Não há informações disponíveis")
	}
}

func main() {

	r := gin.Default()
	r.GET("/", paginaRaiz)
	r.GET("/empregado/:id", acharEmpregado)
	r.Run()
}
