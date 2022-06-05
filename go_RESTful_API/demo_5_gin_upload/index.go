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

		file, _ := ctx.FormFile("file")
		extension := filepath.Ext(file.Filename)

		ctx.SaveUploadedFile(file, fmt.Sprintf("%s/uploaded/%d%s", runningDir, count, extension))
		ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	route.Run(":8080")
}
