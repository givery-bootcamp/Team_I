package controllers

import (
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostCommentInput struct {
	PostId int    `json:"post_id"`
	Body   string `json:"body"`
}

func PostComment(ctx *gin.Context, usecase *usecases.CreateCommentUsecase) {
	var postCommentInput PostCommentInput
	if err := ctx.ShouldBindJSON(&postCommentInput); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	userInfoAny, exists := ctx.Get("userInfo")
	if !exists {
		handleError(ctx, http.StatusBadRequest, ErrUserInfoNotFound)
		return
	}
	userInfo := userInfoAny.(map[string]any)
	userId := userInfo["Id"].(int)
	result, err := usecase.Execute(userId, postCommentInput.PostId, postCommentInput.Body)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

type PutCommentInput struct {
	Body string `json:"body"`
}

func PutComment(ctx *gin.Context, usecase usecases.IUpdateCommentUsecase) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}

	var putCommentInput PutCommentInput
	if err := ctx.ShouldBindJSON(&putCommentInput); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := usecase.Execute(id, putCommentInput.Body)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
