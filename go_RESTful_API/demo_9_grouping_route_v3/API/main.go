package api

import (
	"github.com/gin-gonic/gin"
)

func Setup(route *gin.Engine) {

	authenAPI := route.Group("/authen")
	{

		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
		authenAPI.GET("/profile", Profile)
	}

	stockAPI := route.Group("/stock")
	{

		stockAPI.GET("/list", List)
		stockAPI.GET("/create", Create)
		stockAPI.GET("/edit", Edit)
	}
}
