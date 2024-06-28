package usecases

import (
	"fmt"
	"myapp/internal/entities"
)

type UpdatePostUsecase struct {
	repository entities.PostRepository
}

func NewUpdatePostUsecase(r entities.PostRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{
		repository: r,
	}
}

var ErrNoPermission = fmt.Errorf("cannot edit other user's post")

func (u *UpdatePostUsecase) Execute(id, userId int, title, body string) (*entities.Post, error) {
	post, err := u.repository.GetPostById(id)
	if err != nil {
		return nil, err
	}
	if userId != post.UserId {
		return nil, ErrNoPermission
	}
	return u.repository.Update(id, title, body)
}
