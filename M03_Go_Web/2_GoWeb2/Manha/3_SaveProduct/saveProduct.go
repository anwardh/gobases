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

var products []produto
var lastID int

func Salvar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req produto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)
		c.JSON(http.StatusCreated, req)
	}
}

func exibirProdutos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func main() {

	r := gin.Default()
	r.GET("/produtos", exibirProdutos)

	pr := r.Group("/produtos")
	pr.POST("/salvar", Salvar())
	r.Run()

}
