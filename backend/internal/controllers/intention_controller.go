package controllers

import (
	"fmt"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IntentionRequest struct {
	Status string `json:"state"`
}

func GetIntention(ctx *gin.Context, usecase *usecases.GetIntentionUsecase) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid Post ID format")
		return
	}

	status := ctx.DefaultQuery("state", "")
	if status == "" {
		ctx.String(http.StatusBadRequest, "Status is required")
		return
	}

	result, err := usecase.Execute(postId, status)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func PostIntention(ctx *gin.Context, usecase *usecases.CreateIntentionUsecase) {
	var intensionRequest IntentionRequest
	if err := ctx.ShouldBindJSON(&intensionRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	userInfoAny, exists := ctx.Get("userInfo")
	if !exists {
		fmt.Println("User info not found")
		handleError(ctx, http.StatusBadRequest, ErrUserInfoNotFound)
		return
	}
	userInfo := userInfoAny.(map[string]any)
	userId := userInfo["Id"].(int)

	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid Post ID format")
		return
	}

	result, err := usecase.Execute(userId, postId, intensionRequest.Status)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
