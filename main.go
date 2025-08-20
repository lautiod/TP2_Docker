package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola Mundo desde Gin 🚀")
	})

	r.Run(":8080") // corre en http://localhost:8080
}
