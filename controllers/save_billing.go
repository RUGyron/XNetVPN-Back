package controllers

import (
	yookassapackage "XNetVPN-Back/services/yookassa"
	"github.com/gin-gonic/gin"
)

func SaveBilling(c *gin.Context) {
	err := yookassapackage.RequestBillingSave(c.Param("email"))
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
}
