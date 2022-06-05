package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0
	route.MaxMultipartMemory = 8 << 20 // 8Mb
	route.POST("/upload", func(ctx *gin.Context) {
		count = count + 1
		username := ctx.PostForm("username")
		token := ctx.PostForm("token")
		file, _ := ctx.FormFile("file")
		extension := filepath.Ext(file.Filename)

		ctx.SaveUploadedFile(file, fmt.Sprintf("%s/uploaded/%d%s", runningDir, count, extension))
		ctx.String(200, fmt.Sprintf("'%s'-[username:%s, token:%s] uploaded!",
			file.Filename, username, token))
	})

	route.Run(":8080")
}
