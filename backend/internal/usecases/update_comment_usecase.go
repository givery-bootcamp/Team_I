package usecases

import (
	"myapp/internal/entities"
)

type IUpdateCommentUsecase interface {
	Execute(id int, body string) (*entities.Comment, error)
}

type UpdateCommentUsecase struct {
	repository entities.CommentRepository
}

func NewUpdateCommentUsecase(r entities.CommentRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		repository: r,
	}
}

func (u *UpdateCommentUsecase) Execute(id int, body string) (*entities.Comment, error) {
	comment := entities.Comment{
		Id:   id,
		Body: body,
	}
	return u.repository.Update(&comment)
}
