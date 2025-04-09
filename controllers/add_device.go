package controllers

import (
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories/repo_devices"
	"XNetVPN-Back/responses"
	"XNetVPN-Back/services/jwt"
	"github.com/gin-gonic/gin"
)

func AddDevice(c *gin.Context) {
	var payload in.AddDevice

	user, err := jwt.GetUserByJWT(c)
	if err != nil || user == nil {
		c.JSON(responses.Unauthorized())
		return
	}

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

	device, err := repo_devices.FindDevice(payload.Identifier)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}
	if device != nil {
		c.JSON(responses.Forbidden("device exist"))
		return
	}

	// Enter game in db
	err = repo_devices.InsertDevice(payload, user.Id)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}

	// Success
	c.JSON(responses.Success())
}
