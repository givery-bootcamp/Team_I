package usecases

import (
	"fmt"
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

var ErrInvalidPostType = fmt.Errorf("invalid post type")

func (u *ListPostUsecase) Execute(page int, limit int, postType string) ([]*entities.Post, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 50
	}
	if postType != "" && postType != "official" && postType != "yamada" {
		return nil, ErrInvalidPostType
	}

	return u.repository.List(page, limit, postType)
}
