package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type ListPostUsecase struct {
	repository entities.PostRepository
}

func NewListPostUsecase(r *repositories.PostRepository) *ListPostUsecase {
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

func NewGetPostByIdUsecase(r *repositories.PostRepository) *GetPostByIdUsecase {
	return &GetPostByIdUsecase{
		repository: r,
	}
}

func (u *GetPostByIdUsecase) Execute(id int) (*entities.Post, error) {
	return u.repository.GetPostById(id)
}
