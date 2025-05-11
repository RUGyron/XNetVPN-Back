package controllers

import (
	"XNetVPN-Back/repositories/yk_events"
	"XNetVPN-Back/responses"
	"XNetVPN-Back/services/yookassa"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveBilling(c *gin.Context) {
	email := c.Query("email")
	fmt.Println("email:", email)

	// send request in yk
	ykId, redirectUri, err := yookassa.RequestBillingSave(email)
	if err != nil || redirectUri == nil || ykId == nil {
		c.JSON(responses.ServerError("failed to process payment"))
		return
	}

	// save request in db
	if err := yk_events.InsertBillingSave(*ykId, email); err != nil {
		c.JSON(responses.ServerError("failed to process billing save"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"uri": redirectUri})
}
