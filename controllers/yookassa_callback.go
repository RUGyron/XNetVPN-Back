package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func YookassaCallback(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	var pretty bytes.Buffer
	if err := json.Indent(&pretty, body, "", "  "); err != nil {
		// если не удалось отформатировать, просто выведем как есть
		fmt.Println("Raw JSON:\n", string(body))
	} else {
		fmt.Println("Formatted JSON:\n", pretty.String())
	}

	c.Status(http.StatusOK)
}
