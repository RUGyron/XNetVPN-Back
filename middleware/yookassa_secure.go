package middleware

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
	"net"
	"sync"
)

var (
	once       sync.Once
	allowedIps []*net.IPNet
)

// GetConfig Singleton Config
func getYKAllowedIps() []*net.IPNet {
	once.Do(func() {
		for _, cidr := range config.Config.YooKassaAllowedIps {
			_, ipnet, err := net.ParseCIDR(cidr)
			if err != nil {
				panic("invalid CIDR in Yookassa list: " + cidr)
			}
			allowedIps = append(allowedIps, ipnet)
		}
	})
	return allowedIps
}

func IsYooKassaIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	for _, network := range getYKAllowedIps() {
		if network.Contains(ip) {
			return true
		}
	}
	return false
}

func YooKassaSecure() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsYooKassaIP(c.GetHeader(config.Config.IpHeader)) {
			c.AbortWithStatusJSON(responses.Forbidden("you have no access"))
			return
		}
		c.Next()
	}
}
