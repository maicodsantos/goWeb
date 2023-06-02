package json

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

func GetAll(c *gin.Context) {
	mockedUsers := []User{{Id: 1, Nome: "Maicon", Sobrenome: "Santos", Email: "maicon@hotmails.com", Idade: 33, Altura: 179, Ativo: true, DataDeCriacao: "12/07/1989"}, {Id: 2, Nome: "Isa", Sobrenome: "Souza", Email: "isa@hotmail.com", Idade: 27, Altura: 160, Ativo: true, DataDeCriacao: "16/03/1994"}}
	var filtered []User
	for _, value := range mockedUsers {
		if value.Nome == c.Query("nome") {
			filtered = append(filtered, value)
		}
	}

	c.JSON(200, gin.H{"users": filtered})
}

func GetById(c *gin.Context) {
	mockedUsers := []User{{Id: 1, Nome: "Maicon", Sobrenome: "Santos", Email: "maicon@hotmail.com", Idade: 33, Altura: 179, Ativo: true, DataDeCriacao: "12/07/1989"}, {Id: 2, Nome: "Isa", Sobrenome: "Souza", Email: "isa@hotmail.com", Idade: 27, Altura: 160, Ativo: true, DataDeCriacao: "16/03/1994"}}

	var user User

	for _, value := range mockedUsers {
		if fmt.Sprint(value.Id) == c.Param("id") {
			user = value
		}
	}

	c.JSON(200, gin.H{"User": user})
}
