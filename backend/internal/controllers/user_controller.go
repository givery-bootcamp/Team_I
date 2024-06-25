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

func GetUser(ctx *gin.Context, usecase *usecases.GetUserUsecase) {

	anyUserInfo, ok := ctx.Get("userInfo")
	if !ok {
		handleError(ctx, http.StatusInternalServerError, errors.New("user info not found"))
		return
	}
	userInfo, ok := anyUserInfo.(map[string]interface{})
	if !ok {
		handleError(ctx, http.StatusInternalServerError, errors.New("user info cannot be converted"))
		return
	}
	userId, ok := userInfo["Id"].(int)
	if !ok {
		handleError(ctx, http.StatusInternalServerError, errors.New("user id not found"))
		return
	}
	user, err := usecase.Execute(userId)
	if err != nil {
		log.Printf("Unexpected error in GetUser: %v", err)
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
