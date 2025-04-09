package controllers

import (
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/responses"
	"XNetVPN-Back/services/jwt"
	"XNetVPN-Back/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateToken(c *gin.Context) {
	// Parse JWT
	parsedToken, err := utils.ParseToken(c)
	if err != nil {
		c.JSON(responses.Unauthorized())
		return
	}
	response := out.Login{Tokens: models.Tokens{RefreshToken: parsedToken}}

	// Find user
	user, err := jwt.GetUserByJWT(c)
	if err != nil {
		c.JSON(responses.Unauthorized())
		return
	}
	response.FillWith(user)

	// Update access token
	err = response.Tokens.UpdateAccessToken(user.Id.Hex())
	if err != nil {
		c.JSON(responses.Unauthorized())
		return
	}

	c.JSON(http.StatusOK, response)
}
