package main

import (
	"github.com/gin-gonic/gin"
	json "github.com/maicodsantos/goWeb/json"
)

func main() {
	router := gin.Default()

	router.GET("ola-eu", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola, Maicon"})
	})
	router.GET("users", json.GetAll)
	router.GET("users/:id", json.GetById)
	router.POST("users", json.Create())

	router.Run()
}
