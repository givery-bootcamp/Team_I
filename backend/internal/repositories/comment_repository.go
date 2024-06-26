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

func (r *CommentRepository) Create(comment *entities.Comment) (*entities.Comment, error) {
	result := r.Conn.Table("comments").Select("UserId", "PostId", "Body").Create(comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return comment, nil
}

func (r *CommentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	err := r.Conn.Model(comment).Select("Body").Updates(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
