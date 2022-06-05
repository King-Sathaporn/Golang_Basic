package api

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "login")
}

func Profile(c *gin.Context) {
	c.String(200, "profile")
}

func Register(c *gin.Context) {
	c.String(200, "register")
}

