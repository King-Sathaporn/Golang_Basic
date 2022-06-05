package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"status":  http.StatusOK,
		})
	})

	route.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			message string
			status  int
		}
		msg.Name = "John"
		msg.message = "Hello World"
		msg.status = http.StatusOK
		c.JSON(200, msg)
	})

	route.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "Hello World",
			"status":  http.StatusOK,
		})
	})

	route.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"message": "Hello World",
			"status":  http.StatusOK,
		})
	})

	route.GET("/someString", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	route.Run(":8080")
}
