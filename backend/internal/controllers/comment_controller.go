package controllers

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context, usecase *usecases.CreateCommentUsecase) {
	var comment entities.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	fmt.Println(comment)
	userInfoAny, exists := ctx.Get("userInfo")
	if !exists {
		handleError(ctx, http.StatusBadRequest, ErrUserInfoNotFound)
		return
	}
	userInfo := userInfoAny.(map[string]any)
	userId := userInfo["Id"].(int)
	result, err := usecase.Execute(userId, comment.PostId, comment.Body)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
