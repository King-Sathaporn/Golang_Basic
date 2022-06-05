package api

import (
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.String(200, "list")
}

func Create(c *gin.Context) {
	c.String(200, "create")
}

func Edit(c *gin.Context) {
	c.String(200, "edit")
}

