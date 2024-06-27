package usecases

import (
	"myapp/internal/entities"
)

type IDeleteCommentUsecase interface {
	Execute(user_id, comment_id int) error
}

type DeleteCommentUsecase struct {
	repository entities.CommentRepository
}

func NewDeleteCommentUsecase(r entities.CommentRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{
		repository: r,
	}
}

func (u *DeleteCommentUsecase) Execute(user_id, comment_id int) error {
	// idでコメントを取得して、存在チェックをする
	comment, err := u.repository.GetById(comment_id)
	if err != nil {
		return &CommentNotFound{}
	}
	// ユーザーIDとコメントのユーザーIDが一致するかチェック
	if user_id != comment.UserId {
		return &NoPermission{}
	}
	err = u.repository.Delete(comment_id)
	return err
}
