package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ListPostUsecase struct {
	repository entities.PostRepository
}

func NewPostUsecase(ctx *gin.Context) *ListPostUsecase {
	r := repositories.NewPostRepository(DB(ctx))
	return &ListPostUsecase{
		repository: r,
	}
}

func (u *ListPostUsecase) Execute() ([]*entities.Post, error) {
	return u.repository.List()
}
