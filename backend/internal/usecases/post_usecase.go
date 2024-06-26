package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type ListPostUsecase struct {
	repository entities.PostRepository
}

func NewListPostUsecase(r *repositories.PostRepository) *ListPostUsecase {
	return &ListPostUsecase{
		repository: r,
	}
}

func (u *ListPostUsecase) Execute() ([]*entities.Post, error) {
	return u.repository.List()
}

type GetPostByIdUsecase struct {
	repository entities.PostRepository
}

func NewGetPostByIdUsecase(r *repositories.PostRepository) *GetPostByIdUsecase {
	return &GetPostByIdUsecase{
		repository: r,
	}
}

func (u *GetPostByIdUsecase) Execute(id int) (*entities.Post, error) {
	return u.repository.GetPostById(id)
}

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
