package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success Success
func Success(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   payload,
	})
}

// BadRequest BadRequest
func BadRequest(c *gin.Context, errorCode int, payload interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"result":     false,
		"error_code": errorCode,
		"msg":        payload,
	})
}

// SystemError SystemError
func SystemError(c *gin.Context, errorCode int, payload interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"result":     false,
		"error_code": errorCode,
		"msg":        payload,
	})
}
