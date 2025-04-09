package controllers

import (
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/repositories/repo_devices"
	"XNetVPN-Back/repositories/repo_subscriptions"
	"XNetVPN-Back/responses"
	"XNetVPN-Back/services/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	var response out.Profile

	user, err := jwt.GetUserByJWT(c)
	if err != nil || user == nil {
		c.JSON(responses.Unauthorized())
		return
	}

	devices, err := repo_devices.FindUserDevices(user.Id)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}

	subscription, err := repo_subscriptions.FindUserSubscription(*user.SubscriptionId)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}

	response.FillWith(user, devices, *subscription)

	c.JSON(http.StatusOK, response)
}
