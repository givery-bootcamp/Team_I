package usecases

import (
	"myapp/internal/entities"
)

type IDeleteCommentUsecase interface {
	Execute(comment_id int) error
}

type DeleteCommentUsecase struct {
	repository entities.CommentRepository
}

func NewDeleteCommentUsecase(r entities.CommentRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{
		repository: r,
	}
}

func (u *DeleteCommentUsecase) Execute(comment_id int) error {
	return u.repository.Delete(comment_id)
}
