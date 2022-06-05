package main

import (
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(200, "login")
}

func profile(c *gin.Context) {
	c.String(200, "profile")
}

func register(c *gin.Context) {
	c.String(200, "register")
}

func list(c *gin.Context) {
	c.String(200, "list")
}

func create(c *gin.Context) {
	c.String(200, "create")
}

func edit(c *gin.Context) {
	c.String(200, "edit")
}

//create group router
func main() {
	route := gin.Default()

	//create authenAPI main route
	authenAPI := route.Group("/authen")
	{
		//create authenAPI sub route
		authenAPI.GET("/login", login)
		authenAPI.GET("/register", register)
		authenAPI.GET("/profile", profile)
	}

	//create stockAPI main route
	stockAPI := route.Group("/stock")
	{
		//create stockAPI sub route
		stockAPI.GET("/list", list)
		stockAPI.GET("/create", create)
		stockAPI.GET("/edit", edit)
	}

	route.Run(":8080")
}
