package controllers

import (
	"errors"
	"log"
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
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, password := json.Name, json.Password
	user, tokenString, err := usecase.Execute(username, password)
	if err != nil {
		if err == usecases.ErrUserNotFound || err == usecases.ErrPasswordIncorrect {
			log.Printf("Error in user authentication: %v", err)
			handleError(ctx, http.StatusBadRequest, errors.New("incorrect username or password"))
		} else {
			log.Printf("Unexpected error in PostSignin: %v", err)
			handleError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.SetSameSite(http.SameSiteStrictMode)
	// ヘッダーにトークンをセット
	ctx.SetCookie("jwt", tokenString, 3600, "/", "", false, true)
	ctx.JSON(http.StatusOK, user)
}
