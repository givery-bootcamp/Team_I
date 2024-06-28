package usecases

import (
	"myapp/internal/entities"
)

type GetPostByIdUsecase struct {
	postRepository    entities.PostRepository
	commentRepository entities.CommentRepository
}

func NewGetPostByIdUsecase(postRepository entities.PostRepository, commentRepository entities.CommentRepository) *GetPostByIdUsecase {
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
