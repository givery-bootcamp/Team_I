package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/usecases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JsonRequestUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

const SECRET_KEY = "secret"

func PostSignin(ctx *gin.Context, usecase *usecases.PostSigninUsecase) {

	var json JsonRequestUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username, password := json.Name, json.Password
	result, err := usecase.Execute(username, password)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else if result != nil {
		// トークンの発行（ヘッダー・ペイロード）
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": result.Name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(SECRET_KEY))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
			return
		}
		// ヘッダーにトークンをセット
		ctx.Header("Authorization", tokenString)
		fmt.Println("tokenString")
		fmt.Println(tokenString)
		ctx.JSON(http.StatusOK, result)
	} else {
		handleError(ctx, http.StatusNotFound, errors.New("not found"))
	}

}
