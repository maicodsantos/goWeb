package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maicodsantos/goWeb/cmd/server/handler"
	"github.com/maicodsantos/goWeb/internal/users"
)

func main() {
	repo := users.NewRepository()     // Criação da instância Repository
	service := users.NewService(repo) // Criação da instância Service
	u := handler.NewUser(service)     // Criação do Controller

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/post", u.Create())
	pr.GET("/get", u.GetAll())
	r.Run()
}
