package web

import "github.com/gin-gonic/gin"

func Run(defaultPort string) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"running": "true",
		})
	})

	r.GET("/v1/:org/:project/:pr", func(c *gin.Context) {
		_ = c.Param("org")
		_ = c.Param("project")
		_ = c.Param("pr")

		c.JSON(200, gin.H{
			"backup": true,
			"cache":  false,
		})

	})

	r.Run(":" + defaultPort)
}
