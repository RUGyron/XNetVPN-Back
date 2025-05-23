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
	ApiKey   string `json:"api_key"`
	SiteUrl  string `json:"site_url"`
	IpHeader string `json:"ip_header"`

	// JWT Auth
	JwtKey               string `json:"jwt_key"`
	JwtAccessExpiration  int    `json:"jwt_access_expiration"`  // minutes
	JwtRefreshExpiration int    `json:"jwt_refresh_expiration"` // minutes

	// Mongo Connect
	MongoUri      string `json:"mongo_uri"`
	MongoDatabase string `json:"mongo_database"`

	// Collections
	MongoCollectionUsers         string `json:"mongo_collection_users"`
	MongoCollectionSubscriptions string `json:"mongo_collection_subscriptions"`
	MongoCollectionDevices       string `json:"mongo_collection_devices"`
	MongoCollectionWgConfigs     string `json:"mongo_collection_wg_configs"`
	MongoCollectionConfig        string `json:"mongo_collection_config"`
	MongoCollectionYKEvents      string `json:"mongo_collection_yk_events"`

	// CORS
	CorsAllowedOrigins   []string `json:"cors_allowed_origins"`
	CorsAllowedMethods   []string `json:"cors_allowed_methods"`
	CorsAllowedHeaders   []string `json:"cors_allowed_headers"`
	CorsExposedHeaders   []string `json:"cors_exposed_headers"`
	CorsMaxAge           int      `json:"cors_max_age"`
	CorsAllowCredentials bool     `json:"cors_allow_credentials"`

	// Timeouts (ms)
	TimeoutMongoConnect     int `json:"timeout_mongo_connect"`
	TimeoutMongoQuery       int `json:"timeout_mongo_query"`
	TimeoutMongoQueryInside int `json:"timeout_mongo_query_inside"`
	TimeoutExternalHttp     int `json:"timeout_external_http"`

	// Wg server
	WgClientUsername string `json:"wg_client_username"`
	WgClientPassword string `json:"wg_client_password"`
	WgServerApiUrl   string `json:"wg_server_api"`

	// YooKassa
	YooKassaApiKey     string   `json:"yoo_kassa_api_key"`
	YooKassaShopId     string   `json:"yoo_kassa_shop_id"`
	YooKassaAllowedIps []string `json:"yoo_kassa_allowed_ips"`
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
	c.SiteUrl = os.Getenv("SITE_URL")
	c.JwtKey = os.Getenv("JWT_KEY")
	c.MongoUri = os.Getenv("MONGO_URI")
	c.WgClientUsername = os.Getenv("WG_CLIENT_USERNAME")
	c.WgClientPassword = os.Getenv("WG_CLIENT_PASSWORD")
	c.WgServerApiUrl = os.Getenv("WG_SERVER_API_URL")
	c.YooKassaApiKey = os.Getenv("YOO_KASSA_API_KEY")
	c.YooKassaShopId = os.Getenv("YOO_KASSA_SHOP_ID")
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
