package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ListPostUsecase struct {
	repository entities.PostRepository
}

func NewListPostUsecase(ctx *gin.Context) *ListPostUsecase {
	r := repositories.NewPostRepository(DB(ctx))
	return &ListPostUsecase{
		repository: r,
	}
}

func (u *ListPostUsecase) Execute() ([]*entities.Post, error) {
	return u.repository.List()
}

type GetPostByIdUsecase struct {
	repository entities.PostRepository
}

func NewGetPostByIdUsecase(ctx *gin.Context) *GetPostByIdUsecase {
	r := repositories.NewPostRepository(DB(ctx))
	return &GetPostByIdUsecase{
		repository: r,
	}
}

func (u *GetPostByIdUsecase) Execute(id int) (*entities.Post, error) {
	return u.repository.GetPostById(id)
}
