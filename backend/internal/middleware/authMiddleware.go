package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret"

func AuthMiddleware(ctx *gin.Context) {
	// Cookieヘッダーからトークンを取得
	tokenString, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}

	// トークンの検証
	payload, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil || !payload.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	claims, ok := payload.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	userInfo := make(map[string]any)
	for key, value := range claims {
		userInfo[key] = value
	}
	ctx.Set("userInfo", userInfo)

	ctx.Next()
}
