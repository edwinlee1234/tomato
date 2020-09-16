package middleware

import (
	"net/http"
	"tomato/pkg/jwt"
	"tomato/pkg/res"

	"github.com/gin-gonic/gin"
)

// Auth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		token, err := c.Cookie(jwt.Key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": res.ErrUnauthorizedCode,
			})
			return
		}

		// 解析token 取得會員的資料
		userID, userName, err := jwt.ParseToken(token)
		if err != nil || userID == "" || userName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": res.ErrUnauthorizedCode,
			})
			return
		}

		// 把值傳到下一層
		c.Set("user_id", userID)
		c.Set("user_name", userName)

		c.Next()
	}
}
