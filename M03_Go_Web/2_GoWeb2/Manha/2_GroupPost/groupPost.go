package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type produto struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

func Salvar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req produto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4
		c.JSON(http.StatusCreated, req)
	}
}

func main() {

	r := gin.Default()
	pr := r.Group("/produtos")
	pr.POST("/salvar", Salvar())
	r.Run()

}
