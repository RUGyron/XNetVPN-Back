package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xnet-vpn/models"
	"xnet-vpn/models/out"
	"xnet-vpn/repositories/repo_users"
	"xnet-vpn/responses"
)

func Register(c *gin.Context) {
	var response out.Login

	// Insert new user
	userId, err := repo_users.InsertNewUser()
	if err != nil || userId == nil {
		c.JSON(responses.ServerError())
		return
	}

	// Find user in db
	user, err := repo_users.FindUserById(*userId)
	if err != nil || user == nil {
		c.JSON(responses.Forbidden())
		return
	}

	response.FillWith(user)

	// Generate JWT tokens
	var tokens = models.Tokens{}
	err = tokens.GenerateTokens(userId.Hex())
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}
	response.Tokens = tokens
	c.JSON(http.StatusOK, response)
}
