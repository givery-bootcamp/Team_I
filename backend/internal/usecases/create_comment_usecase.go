package usecases

import (
	"myapp/internal/entities"

	"github.com/go-sql-driver/mysql"
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
	result, err := u.repository.Create(&comment)
	me, ok := err.(*mysql.MySQLError)

	if !ok {
		return result, err
	}
	if me.Number == 1452 {
		return nil, &PostIdNotFound{}
	}
	return result, err
}
