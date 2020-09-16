package handler

import (
	"fmt"
	"tomato/pkg/jwt"
	"tomato/pkg/res"

	"github.com/gin-gonic/gin"
)

// GetUserInfo GetUserInfo
func GetUserInfo(c *gin.Context) {
	token, err := c.Cookie(jwt.Key)
	if err != nil {
		res.Success(c, gin.H{
			"is_login":  false,
			"user_name": "",
			"user_id":   "",
		})
		return
	}

	id, name, err := jwt.ParseToken(token)
	fmt.Println(err)
	if err != nil {
		res.Success(c, gin.H{
			"is_login":  false,
			"user_name": "",
			"user_id":   "",
		})
		return
	}

	res.Success(c, gin.H{
		"is_login":  true,
		"user_name": name,
		"user_id":   id,
	})
}
