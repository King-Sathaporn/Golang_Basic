package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Static("/assets", "./public")//! Should not be set RootPath to public.

	route.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	route.Run(":8080")
}
