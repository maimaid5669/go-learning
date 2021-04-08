package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//每次请求的时候，在终端输出请求路径（http.Request.URL）
		c.String(200, c.FullPath())
		address, _ := os.Getwd()
		file, _ := os.Open(address)
		names, _ := file.Readdirnames(100)
		//每次请求的时候，在终端输出当前目录的文件列表
		c.String(200, "["+strings.Join(names, ",")+"]")
		//每次请求的时候，从请求路径中读出文件名，并且判断是否存在
		filename := c.Query("filename")
		if filename != "" {
			for _, v := range names {
				if v == filename {
					c.String(200, filename+"is here!\n")
					textFile, _ := os.Open(filename)
					for {
						//每次请求的时候，将内容写入返回中
						var b []byte = make([]byte, 2*1024)
						n, _ := textFile.Read(b)
						if n != 0 {
							c.String(200, string(b))
						} else {
							break
						}
					}
					break
				}
				c.String(404, "There is no file named"+filename)
			}
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
