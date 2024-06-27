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
	commentIdString := ctx.Param("id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}

	var putCommentInput PutCommentInput
	if err := ctx.ShouldBindJSON(&putCommentInput); err != nil {
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

	result, err := usecase.Execute(userId, commentId, putCommentInput.Body)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func DeleteComment(ctx *gin.Context, usecase usecases.IDeleteCommentUsecase) {
	commentIdString := ctx.Param("id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	userInfoAny, exists := ctx.Get("userInfo")
	if !exists {
		handleError(ctx, http.StatusBadRequest, ErrUserInfoNotFound)
		return
	}
	userInfo := userInfoAny.(map[string]any)
	userId := userInfo["Id"].(int)

	err = usecase.Execute(userId, commentId)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
