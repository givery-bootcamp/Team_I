package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	List() ([]*entities.Post, error)
}
