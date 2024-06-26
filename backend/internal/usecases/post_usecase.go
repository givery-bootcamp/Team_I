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
	postRepository    entities.PostRepository
	commentRepository entities.CommentRepository
}

func NewGetPostByIdUsecase(postRepository *repositories.PostRepository, commentRepository *repositories.CommentRepository) *GetPostByIdUsecase {
	return &GetPostByIdUsecase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
	}
}

func (u *GetPostByIdUsecase) Execute(user_id int, include_comments bool) (*entities.Post, error) {
	// ユーザーIDから投稿を取得
	post, error := u.postRepository.GetPostById(user_id)
	if error != nil {
		return nil, error
	}

	if include_comments {
		// 投稿からコメントを取得
		comments, error := u.commentRepository.GetListByPostId(post.Id)
		if error != nil {
			return nil, error
		}

		// コメントを投稿に追加
		post.Comments = comments
	}

	return post, nil
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
