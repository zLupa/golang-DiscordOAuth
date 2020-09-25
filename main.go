package main

import (
	"discordOAuth/routes"
	"github.com/gin-gonic/gin"
	"os"
)

type ErrorStruct struct {
	StatusCode int
	Message string
	Error error
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	api := r.Group("/api")
	{
		api.GET("/discordAuth", func(c *gin.Context) {
			c.Redirect(302, os.Getenv("DISCORD_OAUTH_URL"))
		})
		api.GET("/callback", routes.Callback)
	}

	r.Run(":3333")
}
