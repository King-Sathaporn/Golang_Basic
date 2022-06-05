package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	//create route for static files
	route.GET("/image", func(c *gin.Context) {
		c.File("static/test.jpg")
	})

	route.GET("/html", func(c *gin.Context) {
		c.File("static/index.html")
	})

	//create route for downloading files
	route.GET("/download", func(c *gin.Context) {

		c.Header("Content-Disposition", "Simulation File Download")
		c.Header("Countent-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+"image.jpg")
		c.Header("Content-Type", "application/octet-stream")

		c.File("static/test.jpg")
	})

	route.Run(":8080")
}
