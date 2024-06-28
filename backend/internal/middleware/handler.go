package middleware

import (
	"myapp/internal/controllers"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	PostRepository      *repositories.PostRepository
	UserRepository      *repositories.UserRepository
	CommentRepository   *repositories.CommentRepository
	IntentionRepository *repositories.IntentionRepository
}

func NewHandler(postRepository *repositories.PostRepository, userRepository *repositories.UserRepository,
	commentRepository *repositories.CommentRepository, intentionRepository *repositories.IntentionRepository) *Handler {
	return &Handler{
		PostRepository:      postRepository,
		UserRepository:      userRepository,
		CommentRepository:   commentRepository,
		IntentionRepository: intentionRepository,
	}
}

func (h *Handler) GetPosts(ctx *gin.Context) {
	usecase := usecases.NewListPostUsecase(h.PostRepository)
	controllers.GetPosts(ctx, usecase)
}

func (h *Handler) GetPostById(ctx *gin.Context) {
	usecase := usecases.NewGetPostByIdUsecase(h.PostRepository, h.CommentRepository)
	controllers.GetPostById(ctx, usecase)
}

func (h *Handler) DeletePost(ctx *gin.Context) {
	usecase := usecases.NewDeletePostUsecase(h.PostRepository)
	controllers.DeletePost(ctx, usecase)
}
func (h *Handler) PostPost(ctx *gin.Context) {
	usecase := usecases.NewCreatePostUsecase(h.PostRepository)
	controllers.PostPost(ctx, usecase)
}

func (h *Handler) PutPostById(ctx *gin.Context) {
	usecase := usecases.NewUpdatePostUsecase(h.PostRepository)
	controllers.PutPostById(ctx, usecase)
}

func (h *Handler) PostSignup(ctx *gin.Context) {
	signupUsecase := usecases.NewPostSignupUsecase(h.UserRepository)
	signinUsecase := usecases.NewPostSigninUsecase(h.UserRepository)
	controllers.PostSignup(ctx, signupUsecase, signinUsecase)
}

func (h *Handler) PostSignin(ctx *gin.Context) {
	usecase := usecases.NewPostSigninUsecase(h.UserRepository)
	controllers.PostSignin(ctx, usecase)
}

func (h *Handler) PostSignout(ctx *gin.Context) {
	usecase := usecases.NewPostSignoutUsecase(h.UserRepository)
	controllers.PostSignout(ctx, usecase)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	usecase := usecases.NewGetUserUsecase(h.UserRepository)
	controllers.GetUser(ctx, usecase)
}

func (h *Handler) PostComment(ctx *gin.Context) {
	usecase := usecases.NewCreateCommentUsecase(h.CommentRepository)
	controllers.PostComment(ctx, usecase)
}

func (h *Handler) PutComment(ctx *gin.Context) {
	usecase := usecases.NewUpdateCommentUsecase(h.CommentRepository)
	controllers.PutComment(ctx, usecase)
}

func (h *Handler) DeleteComment(ctx *gin.Context) {
	usecase := usecases.NewDeleteCommentUsecase(h.CommentRepository)
	controllers.DeleteComment(ctx, usecase)
}

func (h *Handler) GetIntention(ctx *gin.Context) {
	usecase := usecases.NewGetIntentionUsecase(h.IntentionRepository, h.UserRepository)
	controllers.GetIntention(ctx, usecase)
}

func (h *Handler) PostIntention(ctx *gin.Context) {
	usecase := usecases.NewCreateIntentionUsecase(h.IntentionRepository, h.UserRepository)
	controllers.PostIntention(ctx, usecase)
}
