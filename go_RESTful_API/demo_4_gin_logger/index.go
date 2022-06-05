package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	//gin.DisableConsoleColor() //disable color output in console

	runningDir, _ := os.Getwd()
	coute := 0

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	//route.Use(gin.Logger()) // standard logger
	// custom logger

	route.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//route.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile"))//Hide the path of the log file

	route.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("Root"))
	})
	route.GET("/error", func(ctx *gin.Context) {
		coute = coute + 1
		errlogfile.WriteString(fmt.Sprintf("Error : %d\n", coute))
		ctx.Data(200, "text/html; charset=utf-8", []byte("error"))
	})
	route.GET("/profile", func(ctx *gin.Context) {
		coute = coute + 1
		accesslogfile.WriteString(fmt.Sprintf("count : %d", coute))
		ctx.Data(200, "text/html; charset=utf-8", []byte("profile"))
	})

	route.Run(":8080")
}
