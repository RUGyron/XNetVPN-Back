package controllers

import (
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories/yk_events"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
)

func YookassaCallback(c *gin.Context) {
	var payload in.YookassaCallback

	// Parse JSON
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(responses.InvalidInputs())
		return
	}

	// Confirm yk event
	switch payload.Object.Metadata.Event {
	case models.YKEventType.Save:
		if err := yk_events.UpdateBillingSave(payload); err != nil {
			c.JSON(responses.Forbidden())
			return
		}
	default:
		c.JSON(responses.InvalidInputs())
		return
	}

	c.JSON(responses.Success())
}
