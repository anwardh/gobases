package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type empregado struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Id       string `form:"id" json:"id"`
	Active   string `form:"active" json:"active" binding:"required"` // campo obrigatório
}

func empregadoHandler(c *gin.Context) {

	var empregado empregado
	erro := c.ShouldBindJSON(&empregado)

	if erro != nil {
		// Se acontecer um erro nesta validação, será retornado um Bar Request e uma Interface com um Erro
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	if empregado.Name != "user1" || empregado.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "não está autorizado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mesagem": "Olá ",
	})
}

func main() {

	r := gin.Default()
	r.POST("/empregado", empregadoHandler)
	r.Run()

}
