package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type Model struct {
	// Base
	ApiKey string `json:"api_key"`

	// CORS
	CorsAllowedOrigins   []string `json:"cors_allowed_origins"`
	CorsAllowedMethods   []string `json:"cors_allowed_methods"`
	CorsAllowedHeaders   []string `json:"cors_allowed_headers"`
	CorsExposedHeaders   []string `json:"cors_exposed_headers"`
	CorsMaxAge           int      `json:"cors_max_age"`
	CorsAllowCredentials bool     `json:"cors_allow_credentials"`
}

func (c *Model) ImportJSON() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return
	}
}

func (c *Model) ImportEnv() {
	c.ApiKey = os.Getenv("API_KEY")
}

var (
	once     sync.Once
	instance *Model
	Config   = GetConfig()
)

// GetConfig Singleton Config
func GetConfig() *Model {
	once.Do(func() {
		instance = new(Model)
		instance.ImportJSON()
		instance.ImportEnv()
	})
	return instance
}
