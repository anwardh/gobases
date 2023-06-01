package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func exibirAluno(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []struct {
		Id   string
		Nome string
	}{
		{
			Id:   "1",
			Nome: "Usuário1",
		},

		{
			Id:   "2",
			Nome: "Usuário2",
		},
	})
}

func exibirProfessor(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []struct {
		Id   string
		Nome string
	}{
		{
			Id:   "1",
			Nome: "Professor1",
		},

		{
			Id:   "2",
			Nome: "Professor2",
		},
	})
}

func exibirSobre(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Sobre": "Wave 2",
		"Curso": "Backend Go",
	})
}

func main() {

	r := gin.Default()

	grupo := r.Group("/bootcamp/go")
	{
		grupo.GET("/alunos", exibirAluno)
		grupo.GET("/professores", exibirProfessor)
		grupo.GET("/sobre", exibirSobre)
	}
	r.Run()
}
