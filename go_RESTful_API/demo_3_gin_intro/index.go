package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("root"))
	})

	router.GET("profile", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("profile"))
	})
	router.GET("login", func(ctx *gin.Context) {
		username, password := ctx.Query("username"), ctx.Query("password")
		ctx.JSON(http.StatusOK, gin.H{"result": "success", "username": username, "password": password}) //gin.H is hashmap
	})

	router.GET("/book/:from/:to/:vehical", handleBookRequest)

	router.POST("/login", func(ctx *gin.Context) {
		var form LoginForm
		if ctx.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "1234" {
				msg := fmt.Sprintf("You are logged in as %s %s", form.Username, form.Password)
				ctx.JSON(200, gin.H{"status": msg})
			} else {
				ctx.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			ctx.JSON(401, gin.H{"status": "under to bind data"})
		}
	})

	router.Run(":8080")
}

func handleBookRequest(ctx *gin.Context) {
	from, to := ctx.Param("from"), ctx.Param("to")
	vehical := ctx.Param("vehical")
	ctx.JSON(http.StatusOK, gin.H{"result": "success", "from": from, "to": to, "vehical": vehical})
}
