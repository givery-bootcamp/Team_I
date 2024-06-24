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

type DeletePostUsecase struct {
	repository entities.PostRepository
}

func NewDeletePostUsecase(r *repositories.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{
		repository: r,
	}
}

func (u *DeletePostUsecase) Execute(id int) (error) {
	return u.repository.DeletePost(id)
}