package main

import (
	"demo9/api"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	api.Setup(route)

	route.Run(":8080")
}
