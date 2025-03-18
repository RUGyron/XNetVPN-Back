package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Unauthorized 401
func Unauthorized(details ...string) (int, gin.H) {
	return http.StatusUnauthorized, gin.H{"message": "unauthorized", "code": 1, "details": details}
}

// InvalidInputs 400
func InvalidInputs(details ...string) (int, gin.H) {
	return http.StatusBadRequest, gin.H{"message": "invalid inputs", "code": 2, "details": details}
}

// ServerError 500
func ServerError(details ...string) (int, gin.H) {
	return http.StatusInternalServerError, gin.H{"message": "server error", "code": -1, "details": details}
}

// Success 200
func Success(details ...string) (int, gin.H) {
	return http.StatusOK, gin.H{"message": "success", "code": 0, "details": details}
}

// Forbidden 403
func Forbidden(details ...string) (int, gin.H) {
	return http.StatusForbidden, gin.H{"message": "success", "code": 3, "details": details}
}
