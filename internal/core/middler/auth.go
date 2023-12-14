package middler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/protect-we-network/server/internal/packages/config"
)

func getToken(c *gin.Context) string {
	token, error := c.Cookie(config.AuthenticationHeader)
	if error != nil || token == "" {
		return ""
	}

	return token
}

func Auth(onlyAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 解析 JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 这里的密钥需要与签发 JWT token 的密钥一致
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 将当前用户的信息添加到上下文中
		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("user", claims["username"])
		c.Next()
	}
}
