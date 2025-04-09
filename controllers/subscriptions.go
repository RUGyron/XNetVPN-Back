package controllers

import (
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/repositories/repo_subscriptions"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Subscriptions(c *gin.Context) {
	var response out.Subscriptions

	// Find subscriptions in db
	subscriptions, err := repo_subscriptions.FindSubscriptions()
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}

	response.FillWith(subscriptions)

	c.JSON(http.StatusOK, response)
}
