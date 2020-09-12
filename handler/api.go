package handler

import (
	"tomato/pkg/res"

	"github.com/gin-gonic/gin"
)

// Ping Ping
func Ping(c *gin.Context) {
	res.Success(c, gin.H{
		"msg": "pong",
	})
}
