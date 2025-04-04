package controllers

import (
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/repositories/repo_users"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Login(c *gin.Context) {
	var response out.Login

	// Parse user ID
	userId, err := primitive.ObjectIDFromHex(c.Param("key"))
	if err != nil {
		c.JSON(responses.InvalidInputs())
		return
	}

	// Find user in db
	user, err := repo_users.FindUserById(userId)
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

	c.JSON(http.StatusOK, response)
}
