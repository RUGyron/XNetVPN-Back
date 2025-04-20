package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func YookassaCallback(c *gin.Context) {
	var body map[string]interface{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	fmt.Printf("YooKassa callback: %+v\n", body)
	c.Status(http.StatusOK)
}
