package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anynines/pr-config/backend"

	"github.com/gin-gonic/gin"
)

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
			org := c.Param("org")
			project := c.Param("project")
			pr := c.Param("pr")

			// fetch json data and validate proper format
			data, err := c.GetRawData()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "Failed to read json data from body",
				})
				return
			}

			if !json.Valid(data) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "Invalid json",
				})
				return
			}

			// store as json string in backend
			jsonStr := string(data)
			b := backend.NewRedisBackend()
			err = b.Write(fmt.Sprintf("%s/%s/%s/%s", "pr-config", org, project, pr), jsonStr)
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"error": fmt.Sprintf("Backend: %s", err.Error()),
				})
				return
			}
		})
	}

	r.Run(":" + defaultPort)
}
