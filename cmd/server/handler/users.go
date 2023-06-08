package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maicodsantos/goWeb/internal/users"
)

type request struct {
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Email         string `json:"email"`
	Idade         int    `json:"idade"`
	Altura        int    `json:"altura"`
	Ativo         bool   `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *User) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		u, err := c.service.Create(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataDeCriacao)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, u)
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Validação do Token
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// VALIDAÇÃO DAS ATRIBUIÇÕES DOS CAMPOS DA REQUEST
		/*
			Se algum dos atributos for vazio, o Update não ocorrerá - aqui estão as Regras de Negócio para op Update de um produto
			Este Controller serve, justamente, para que os dados coletados na requisição não sejam, diretamente, armazendos
		no Banco de Dados */

		// Validação da Vinculação dos parâmetros para a Estrutura Request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validação do Nome do Usuario
		if req.Nome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do usuário é obrigatório."})
			return
		}

		// Validação do Sobrenome do Usuario
		if req.Sobrenome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O sobrenome do usuário é obrigatório."})
			return
		}

		// Validação da Email do Usuario
		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O email é obrigatório."})
			return
		}

		// Validação da Idade do Usuario
		if req.Idade == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A idade é obrigatória"})
			return
		}

		// Quando estiver 'OK', será chamado o método Update, do Sservice

		u, err := c.service.Update(int(id), req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataDeCriacao)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return // Retorno do erro do Service
		}
		ctx.JSON(http.StatusOK, u) // Retorno "OK" do Service

	}
}
