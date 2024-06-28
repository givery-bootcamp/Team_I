package services

import (
	"log"
	"myapp/internal/controllers"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSignup(ctx *gin.Context, postSignup *usecases.PostSignupUsecase, postSignin *usecases.PostSigninUsecase) {

	var json controllers.JsonRequestUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, password := json.Name, json.Password
	_, err := postSignup.Execute(username, password)
	if err != nil {
		if err == usecases.ErrUnexpected {
			log.Printf("Unexpected error in PostSignup: %v", err)
			handleError(ctx, http.StatusInternalServerError, err)
			return
		}
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	user, tokenString, err := postSignin.Execute(username, password)
	if err != nil {
		if err == usecases.ErrUserNotFound || err == usecases.ErrPasswordIncorrect {
			log.Printf("Error in user authentication: %v", err)
			handleError(ctx, http.StatusUnauthorized, err)
		} else {
			log.Printf("Unexpected error in PostSignin: %v", err)
			handleError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.SetSameSite(http.SameSiteStrictMode)
	// ヘッダーにトークンをセット
	ctx.SetCookie("jwt", tokenString, -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, user)
}
