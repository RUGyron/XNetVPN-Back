package controllers

import (
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/models/out"
	"XNetVPN-Back/repositories/repo_configs"
	"XNetVPN-Back/repositories/repo_devices"
	"XNetVPN-Back/responses"
	"XNetVPN-Back/services/jwt"
	"XNetVPN-Back/services/wg_api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWgConfig(c *gin.Context) {
	var payload in.Config
	var response out.Config

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

	device, err := repo_devices.FindDeviceById(payload.DeviceId, user.Id)
	if err != nil {
		c.JSON(responses.ServerError())
		return
	}
	if device == nil {
		c.JSON(responses.Forbidden("no device"))
		return
	}

	if device.ConfigId == nil {
		wgConfig, err := wg_api.CreateWgConfig()
		if err != nil {
			c.JSON(responses.ServerError())
			return
		}
		configId, err := repo_configs.InsertConfiWg(wgConfig)
		if err != nil {
			c.JSON(responses.ServerError())
			return
		}
		err = repo_devices.UpdateDeviceConfig(device.Id, configId)
		if err != nil {
			c.JSON(responses.ServerError())
			return
		}
		device.ConfigId = &configId
	}
	config, err := repo_configs.FindConfigById(device.ConfigId)
	if err != nil || config == nil {
		c.JSON(responses.ServerError())
		return
	}
	response.FillWith(*config)

	c.JSON(http.StatusOK, response)
}
