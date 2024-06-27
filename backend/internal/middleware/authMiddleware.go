package middleware

import (
	"myapp/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	// Cookieヘッダーからトークンを取得
	tokenString, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Cookie jwt could not be found"})
	}

	// トークンの検証
	payload, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil || !payload.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWT is invalid"})
		ctx.Abort()
		return
	}

	claims, ok := payload.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWT claims could not be read"})
		ctx.Abort()
		return
	}
	userInfo := make(map[string]any)
	for key, value := range claims {
		if key == "Id" {
			userInfo[key] = int(value.(float64))
		} else {
			userInfo[key] = value
		}
	}
	ctx.Set("userInfo", userInfo)
	ctx.Next()
}
