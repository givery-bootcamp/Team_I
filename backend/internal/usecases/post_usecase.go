package usecases

import (
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
	postRepository entities.PostRepository
	commentRepository entities.CommentRepository
}

func NewGetPostByIdUsecase(postRepository *repositories.PostRepository, commentRepository *repositories.CommentRepository) *GetPostByIdUsecase {
	return &GetPostByIdUsecase{
		postRepository: postRepository,
		commentRepository: commentRepository,
	}
}

func (u *GetPostByIdUsecase) Execute(id int) (*entities.Post, error) {
	return u.postRepository.GetPostById(id)
}
