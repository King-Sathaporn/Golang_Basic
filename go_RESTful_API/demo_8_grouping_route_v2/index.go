package main

import (
	"demo8/api"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	//create authenAPI main route
	authenAPI := route.Group("/authen")
	{
		//create authenAPI sub route
		authenAPI.GET("/login", api.Login)
		authenAPI.GET("/register", api.Register)
		authenAPI.GET("/profile", api.Profile)
	}

	//create stockAPI main route
	stockAPI := route.Group("/stock")
	{
		//create stockAPI sub route
		stockAPI.GET("/list", api.List)
		stockAPI.GET("/create", api.Create)
		stockAPI.GET("/edit", api.Edit)
	}

	route.Run(":8080")
}
