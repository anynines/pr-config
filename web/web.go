package web

import "github.com/gin-gonic/gin"

func Run(defaultPort string, username string, password string) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"running": "true",
		})
	})

	authorized := r.Group("/v1", gin.BasicAuth(gin.Accounts{
		username: password,
	}))

	{
		authorized.GET(":org/:project/:pr", func(c *gin.Context) {
			_ = c.Param("org")
			_ = c.Param("project")
			_ = c.Param("pr")

			c.JSON(200, gin.H{
				"backup": true,
				"cache":  false,
			})
		})

		authorized.POST(":org/:project/:pr", func(c *gin.Context) {
		})
	}

	r.Run(":" + defaultPort)
}
