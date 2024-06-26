package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type CommentRepository struct {
	Conn *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) *CommentRepository {
	return &CommentRepository{
		Conn: conn,
	}
}

func (r *CommentRepository) GetListByPostId(post_id int) ([]*entities.Comment, error) {
	var comments []*entities.Comment
	if err := r.Conn.Table("comments").Select("comments.id, comments.user_id, comments.post_id, users.name as username, comments.body, comments.created_at, comments.updated_at").Joins("JOIN users ON comments.user_id = users.id").Where("comments.post_id = ?", post_id).Order("comments.id ASC").Scan(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
