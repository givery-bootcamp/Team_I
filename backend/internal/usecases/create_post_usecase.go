package usecases

import (
	"myapp/internal/entities"
)

type CreatePostUsecase struct {
	repository entities.PostRepository
}

func NewCreatePostUsecase(r entities.PostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		repository: r,
	}
}

func (u *CreatePostUsecase) Execute(userId int, title, body string) (*entities.PostForInsert, error) {
	return u.repository.Create(userId, title, body)
}
