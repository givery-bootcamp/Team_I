package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type DeletePostUsecase struct {
	repository entities.PostRepository
}

func NewDeletePostUsecase(r *repositories.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{
		repository: r,
	}
}

func (u *DeletePostUsecase) Execute(user_id, post_id int) error {
	// ポストidからポストを取得します
	post, err := u.repository.GetPostById(post_id)
	if err != nil {
		return err
	}
	// ポストのuser_idとユーザidを比較します
	// 一致しない場合はエラーを返します
	if post.UserId != user_id {
		return errors.New("user_id does not match")
	}

	// 一致する場合はポストを削除します
	return u.repository.DeletePost(post_id)
}
