package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostUsecase struct {
	repository interfaces.PostRepository
}

func NewPostUsecase(r interfaces.PostRepository) *PostUsecase {
	return &PostUsecase{
		repository: r,
	}
}

func (u *PostUsecase) List() ([]*entities.Post, error) {
	return u.repository.List()
}
