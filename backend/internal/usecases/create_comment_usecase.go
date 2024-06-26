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

func (u *CreateCommentUsecase) Execute(userId, postId int, body string) (*entities.Comment, error) {
	comment := entities.Comment{
		UserId: userId,
		PostId: postId,
		Body:   body,
	}
	return u.repository.Create(&comment)
}
