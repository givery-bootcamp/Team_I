package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context, usecase *usecases.ListPostUsecase) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	result, err := usecase.Execute(page, limit)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else if result != nil {
		ctx.JSON(http.StatusOK, result)
	} else {
		handleError(ctx, http.StatusNotFound, errors.New("not found"))
	}
}

func GetPostById(ctx *gin.Context, usecase *usecases.GetPostByIdUsecase) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	// ctx.String(http.StatusOK, greetings[id])

	include_comments := true
	result, err := usecase.Execute(id, include_comments)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else if result != nil {
		ctx.JSON(http.StatusOK, result)
	} else {
		handleError(ctx, http.StatusNotFound, errors.New("not found"))
	}
}

func DeletePost(ctx *gin.Context, usecase *usecases.DeletePostUsecase) {
	idString := ctx.Param("id")
	post_id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	// ctx.String(http.StatusOK, greetings[id])

	// ユーザidを取得します
	userInfo, is_exists := ctx.Get("userInfo")
	if !is_exists {
		handleError(ctx, http.StatusInternalServerError, errors.New("userInfo does not exist"))
		return
	}

	// userInfoをマップとしてキャスト
	userInfoMap, ok := userInfo.(map[string]interface{})
	if !ok {
		handleError(ctx, http.StatusInternalServerError, errors.New("userInfo is not a map"))
		return
	}
	// userInfoからuser_idを取得します
	user_id, is_exists := userInfoMap["Id"]
	if !is_exists {
		handleError(ctx, http.StatusInternalServerError, errors.New("user_id does not exist"))
		return
	}

	// ポストを削除します
	err = usecase.Execute(user_id.(int), post_id)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "success"})
	}
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ErrUserInfoNotFound = fmt.Errorf("user info not found")

func PostPost(ctx *gin.Context, usecase *usecases.CreatePostUsecase) {
	var post PostRequest
	if err := ctx.ShouldBindJSON(&post); err != nil {
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
	result, err := usecase.Execute(userId, post.Title, post.Content)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func PutPostById(ctx *gin.Context, usecase *usecases.UpdatePostUsecase) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}

	var post PostRequest
	if err := ctx.ShouldBindJSON(&post); err != nil {
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

	result, err := usecase.Execute(id, userId, post.Title, post.Content)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
