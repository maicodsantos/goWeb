package main

import "github.com/gin-gonic/gin"

type User struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Email         string `json:"email"`
	Idade         int    `json:"idade"`
	Altura        int    `json:"altura"`
	Ativo         bool   `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}

func main() {
	users := []User{{Id: 1, Nome: "Maicon", Sobrenome: "Santos", Email: "maicon@gmail.com", Idade: 33, Altura: 179, Ativo: true, DataDeCriacao: "12/07/1989"}, {Id: 2, Nome: "Isa", Sobrenome: "Santos", Email: "isa@gmail.com", Idade: 27, Altura: 160, Ativo: true, DataDeCriacao: "12/01/2000"}}
	router := gin.Default()

	router.GET("ola-eu", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola, Maicon"})
	})
	router.GET("users", func(c *gin.Context) {
		c.JSON(200, gin.H{"users": users})
	})

	router.Run()
}
