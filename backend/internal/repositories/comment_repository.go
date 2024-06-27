package repositories

import (
	"myapp/internal/entities"
	"time"

	"github.com/go-sql-driver/mysql"
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

func (r *CommentRepository) GetById(comment_id int) (*entities.Comment, error) {
	var comment entities.Comment
	if err := r.Conn.Table("comments").Select("id, user_id, post_id, body, created_at, updated_at").Where("id = ? AND deleted_at IS NULL", comment_id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetListByPostId(post_id int) ([]*entities.Comment, error) {
	var comments []*entities.Comment
	if err := r.Conn.Table("comments").Select("comments.id, comments.user_id, comments.post_id, users.name as username, comments.body, comments.created_at, comments.updated_at").Where("comments.deleted_at IS NULL").Joins("JOIN users ON comments.user_id = users.id").Where("comments.post_id = ?", post_id).Order("comments.id ASC").Scan(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) Create(comment *entities.Comment) (*entities.Comment, error) {
	result := r.Conn.Table("comments").Select("UserId", "PostId", "Body").Create(comment)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, result.Error
		}
		if me.Number == FOREIGN_KEY_CONSTRAINTS_ERROR_NUMBER {
			return nil, &PostIdNotFound{}
		}
		return nil, result.Error
	}

	return comment, nil
}

func (r *CommentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	if err := r.Conn.Model(comment).Select("Body").Updates(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) Delete(comment_id int) error {
	t := time.Now()
	if err := r.Conn.Table("comments").Where("id = ? AND deleted_at IS NULL", comment_id).Update("deleted_at", t).Error; err != nil {
		return err
	}
	return nil
}
