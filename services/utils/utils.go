package utils

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/db"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
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

func InitValidator() {
	once.Do(func() {
		localValidator = validator.New()
		for ruleName, regexPattern := range models.RegexRules {
			re := regexp.MustCompile(regexPattern)
			err := localValidator.RegisterValidation(ruleName, func(fl validator.FieldLevel) bool {
				return re.MatchString(fl.Field().String())
			})
			if err != nil {
				panic(err)
			}
		}
	})
}

func CalculateSubscriptionSwitch(currentProduct *db.Product, targetProduct db.Product, startDate time.Time, currentCredit float64) (*models.ProductSwitchResult, error) {
	// new: apply new
	if currentProduct == nil {
		if currentCredit >= targetProduct.Price {
			return &models.ProductSwitchResult{AmountToPay: 0, NewCredit: currentCredit - targetProduct.Price, ApplyNow: true}, nil
		}
		return &models.ProductSwitchResult{AmountToPay: targetProduct.Price - currentCredit, NewCredit: 0, ApplyNow: true}, nil
	}

	// care not same product
	if targetProduct.Rank == currentProduct.Rank {
		return nil, errors.New("target product is the same")
	}

	// downgrade: apply after current
	if targetProduct.Rank <= currentProduct.Rank {
		if currentCredit >= targetProduct.Price {
			return &models.ProductSwitchResult{AmountToPay: 0, NewCredit: currentCredit - targetProduct.Price, ApplyNow: false}, nil
		}
		return &models.ProductSwitchResult{AmountToPay: targetProduct.Price - currentCredit, NewCredit: 0, ApplyNow: false}, nil
	}

	// upgrade: cancel current, apply new

	// calculate used price
	usedDays := time.Now().UTC().Sub(startDate).Hours() / 24
	var totalDays float64
	if currentProduct.Annual {
		totalDays = 365.0
	} else {
		totalDays = 30.0
	}
	usedPrice := (currentProduct.Price / totalDays) * usedDays
	remainingCredit := currentProduct.Price - usedPrice + currentCredit

	if remainingCredit >= targetProduct.Price {
		return &models.ProductSwitchResult{AmountToPay: 0, NewCredit: remainingCredit - targetProduct.Price, ApplyNow: true}, nil
	}
	return &models.ProductSwitchResult{AmountToPay: targetProduct.Price - remainingCredit, NewCredit: 0, ApplyNow: true}, nil
}
