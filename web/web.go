package web

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"running": "true",
		})
	})
	r.Run()
}
