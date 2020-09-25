package routes

import (
	"discordOAuth/requests"
	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {

	codeAuth, stringExists := c.GetQuery("code")

	if !stringExists {
		c.JSON(500, gin.H{"message": "Query string code is missing."})
		return
	}

	r, err := requests.GetToken(codeAuth)

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Message})
		return
	}

	user, err := requests.GetUserInfo(r.AccessToken, r.TokenType)

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Message})
		return
	}

	resp, err := requests.SendNewUserMessage(user)

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Message})
		return
	}



	c.JSON(200, resp)
}

