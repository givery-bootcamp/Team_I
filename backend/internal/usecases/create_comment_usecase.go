package usecases

import (
	"myapp/internal/entities"
)

type CreateCommentUsecase struct {
	repository entities.CommentRepository
}

func NewCreateCommentUsecase(r entities.CommentRepository) *CreateCommentUsecase {
	return &CreateCommentUsecase{
		repository: r,
	}
}

func (u *CreateCommentUsecase) Execute(postId, userId int, body string) (*entities.Comment, error) {
	return u.repository.Create(postId, userId, body)
}
