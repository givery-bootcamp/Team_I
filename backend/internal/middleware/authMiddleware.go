package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret"

func AuthMiddleware(ctx *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	authorizationHeader := ctx.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// トークンの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
