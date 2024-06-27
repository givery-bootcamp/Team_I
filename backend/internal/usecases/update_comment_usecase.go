package usecases

import (
	"myapp/internal/entities"
)

type IUpdateCommentUsecase interface {
	Execute(user_id, comment_id int, body string) (*entities.Comment, error)
}

type UpdateCommentUsecase struct {
	repository entities.CommentRepository
}

func NewUpdateCommentUsecase(r entities.CommentRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		repository: r,
	}
}

func (u *UpdateCommentUsecase) Execute(user_id, comment_id int, body string) (*entities.Comment, error) {
	// idでコメントを取得して、存在チェックをする
	comment, err := u.repository.GetById(comment_id)
	if err != nil {
		return nil, &CommentNotFound{}
	}
	// ユーザーIDとコメントのユーザーIDが一致するかチェック
	if user_id != comment.UserId {
		return nil, &NoPermission{}
	}
	comment.Body = body
	result, err := u.repository.Update(comment)
	return result, err
}
