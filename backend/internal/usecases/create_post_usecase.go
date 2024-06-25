package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type CreatePostUsecase struct {
	repository entities.PostRepository
}

func NewCreatePostUsecase(r *repositories.PostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		repository: r,
	}
}

func (u *CreatePostUsecase) Execute(userId int, title, body string) (*entities.PostForInsert, error) {
	return u.repository.Create(userId, title, body)
}
