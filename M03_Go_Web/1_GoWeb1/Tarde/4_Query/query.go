package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	// Querystring - o que vem na URL -> ? e &
	nome := c.Query("nome") // Retorna o valor da consulta de URL com chave, se existir, caso contrário, retorna uma string vazia
	sobrenome := c.Query("sobrenome")
	log.Println(nome, sobrenome) // Apresenta no console os valores de nome e sobrenome informados na URL
	c.JSON(http.StatusOK, gin.H{
		"mesagem": "Olá " + nome + " " + sobrenome,
	})
}

func helloHandler(c *gin.Context) {
	// URL Param (o que vem depois de ":" -> :id)
	// retorna o valor do parâmetro de URL
	id := c.Param("id")
	_, erro := strconv.Atoi(id)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"menssagem": "o id informado não é um número",
		})
		log.Println(erro)
		return
	}
}

func main() {

	r := gin.Default()
	r.GET("/", rootHandler)

	r.GET("/ola/:id", helloHandler)

	r.Run()

}
