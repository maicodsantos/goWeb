package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Usuarios []User `json:"usuarios"`
}

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
	jsonFile, err := ioutil.ReadFile("./json/usuarios.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var usuarios Users
	json.Unmarshal(jsonFile, &usuarios)

	var usuariosFiltrados []User
	var param = c.Query("nome")
	fmt.Println(param)
	if param != "" {
		for _, usuario := range usuarios.Usuarios {
			if usuario.Nome == param {
				usuariosFiltrados = append(usuariosFiltrados, usuario)
			}

		}
	}

	c.JSON(200, gin.H{"usuarios": usuariosFiltrados})

}

func GetById(c *gin.Context) {
	jsonFile, err := ioutil.ReadFile("./json/usuarios.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var usuarios Users
	json.Unmarshal(jsonFile, &usuarios)

	var usuariosFiltrados []User
	param, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		for _, usuario := range usuarios.Usuarios {
			if usuario.Id == param {
				usuariosFiltrados = append(usuariosFiltrados, usuario)
			}

		}
	}

	if usuariosFiltrados == nil {
		c.JSON(http.StatusNotFound, "Não há informações disponíveis")
	} else {
		c.JSON(200, gin.H{"usuarios": usuariosFiltrados})
	}

}
