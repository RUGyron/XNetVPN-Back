package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
	"xnet-vpn-back/config"
)

var (
	once           sync.Once
	localValidator *validator.Validate
)

func Map[IN any, OUT any](array []IN, fn func(t IN) OUT) []OUT {
	result := make([]OUT, len(array))
	for i, item := range array {
		result[i] = fn(item)
	}
	return result
}

func Contains[T int | string](array []T, value T) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

// ParseToken "Bearer <token>" -> "<token>"
func ParseToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", errors.New("bearer token not in proper format")
	}

	splitToken := strings.SplitN(bearerToken, " ", 2)
	if len(splitToken) != 2 {
		return "", errors.New("bearer token not in proper format")
	}

	return strings.TrimSpace(splitToken[1]), nil
}

func GetIP(c *gin.Context, key string) string {
	header := c.GetHeader(key)
	fwdIPs := strings.Split(header, ",")

	if len(fwdIPs) > 0 {
		return strings.TrimSpace(fwdIPs[0])
	}
	return ""
}

// ValidateStruct Template for fast usage to validate IN controllers' payload
func ValidateStruct[IN interface{}](payload IN) (bool, []string) {
	var reasons []string
	err := localValidator.Struct(payload)

	if err != nil {
		t := reflect.TypeOf(payload)
		for _, validationErr := range err.(validator.ValidationErrors) {
			var rule string
			field, _ := t.FieldByName(validationErr.StructField())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = validationErr.StructField()
			}
			if validationErr.Param() != "" {
				rule = fmt.Sprintf(" (%v)", validationErr.Param())
			}
			reasons = append(reasons, fmt.Sprintf("%s violated \"%s\" validation%v", jsonTag, validationErr.Tag(), rule))
		}
		return false, reasons
	}
	return true, reasons
}

// GetUserIdFromToken "<token>" -> "userId" or error (validated by jwt)
func GetUserIdFromToken(token string) (*string, error) {
	var parsedToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return nil, errors.New("invalid userId in token")
	}

	return &userId, nil
}
