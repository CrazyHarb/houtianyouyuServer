package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		content, err := os.ReadFile("./website/index.html")
		if err != nil {
			c.String(404, "?")
		} else {
			c.Data(200, "text/html; charset=utf-8", content)
		}
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
