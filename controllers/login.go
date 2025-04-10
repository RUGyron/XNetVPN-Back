package controllers

import (
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/db"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/repositories/repo_devices"
	"XNetVPN-Back/repositories/repo_subscriptions"
	"XNetVPN-Back/repositories/repo_users"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Login(c *gin.Context) {
	var payload in.Login
	var response out.Login
	var userId primitive.ObjectID
	var subscription *db.Subscription
	var err error

	// Parse JSON
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(responses.InvalidInputs())
		return
	}

	// Validate payload
	if ok, errs := payload.Validate(); !ok {
		c.JSON(responses.InvalidInputs(errs...))
		return
	}

	if payload.Key == nil || len(*payload.Key) == 0 {
		idPtr, err := repo_users.InsertNewUser()
		if err != nil || idPtr == nil {
			c.JSON(responses.ServerError())
			return
		}
		userId = *idPtr
	} else {
		userId, err = primitive.ObjectIDFromHex(*payload.Key)
		if err != nil {
			c.JSON(responses.InvalidInputs())
			return
		}
	}

	// Find user in db
	user, err := repo_users.FindUserById(userId)
	if err != nil || user == nil {
		c.JSON(responses.Forbidden())
		return
	}

	devices, err := repo_devices.FindUserDevices(user.Id)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}

	if user.SubscriptionId != nil {
		subscription, err = repo_subscriptions.FindUserSubscription(*user.SubscriptionId)
		if err != nil {
			c.JSON(responses.ServerError())
			return
		}
	}

	response.User.FillWith(user, devices, subscription)

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
