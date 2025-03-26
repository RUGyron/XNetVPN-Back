package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"xnet-vpn/models"
	"xnet-vpn/models/db"
	"xnet-vpn/repositories/repo_users"
	"xnet-vpn/responses"
	"xnet-vpn/services/utils"
)

// AccessBearerRequired Auth required
func AccessBearerRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.ParseToken(c)
		if err != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		jwtTokens := models.Tokens{AccessToken: token}
		if jwtTokens.ValidateAccessToken() != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		stringUserId, err := utils.GetUserIdFromToken(jwtTokens.AccessToken)
		if err != nil || stringUserId == nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		userId, err := primitive.ObjectIDFromHex(*stringUserId)
		if err != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		user, err := repo_users.FindUserById(userId)
		if err != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

// RefreshBearerRequired auth via refresh required
func RefreshBearerRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.ParseToken(c)
		if err != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		jwtTokens := models.Tokens{RefreshToken: token}
		if jwtTokens.ValidateRefreshToken() != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		stringUserId, err := utils.GetUserIdFromToken(jwtTokens.RefreshToken)
		if err != nil || stringUserId == nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		userId, err := primitive.ObjectIDFromHex(*stringUserId)
		if err != nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		user, err := repo_users.FindUserById(userId)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(responses.Unauthorized())
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

// GetUserByJWT Requires Access / Refresh Token Required middleware
func GetUserByJWT(c *gin.Context) (*db.User, error) {
	user, exists := c.Get("user")
	if !exists || user == nil {
		return nil, errors.New("unauthorized")
	}
	typedUser, ok := user.(*db.User)
	if !ok {
		return nil, errors.New("failed to parse user")
	}
	return typedUser, nil
}
