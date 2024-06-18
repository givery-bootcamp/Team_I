package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonRequestUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PostSignin(ctx *gin.Context, usecase *usecases.PostSigninUsecase) {

	var json JsonRequestUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username, password := json.Name, json.Password
	user, tokenString, err := usecase.Execute(username, password)
	if err != nil {
		if err == usecases.ErrUserNotFound || err == usecases.ErrPasswordIncorrect {
			handleError(ctx, http.StatusBadRequest, errors.New("incorrect username or password"))
		} else {
			handleError(ctx, http.StatusInternalServerError, err)
		}
		return
	}
	ctx.SetSameSite(http.SameSiteNoneMode)
	// ヘッダーにトークンをセット
	ctx.SetCookie("jwt", tokenString, 3600, "/", "localhost:3000", false, true)
	fmt.Println("tokenString")
	fmt.Println(tokenString)
	ctx.JSON(http.StatusOK, user)

}
