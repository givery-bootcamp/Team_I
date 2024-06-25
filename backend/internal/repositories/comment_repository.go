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
	if err := r.Conn.Table("comments").Select("comments.id, comments.user_id, comments.post_id, users.name as username, comments.body, comments.created_at, comments.updated_at").Joins("JOIN users ON comments.user_id = users.id").Order("comments.id DESC").Scan(&comments).Error; err != nil {
		return nil, err
	}
	// return convertPostRepositoryModelToEntity(posts), nil
	return comments, nil
}

// func convertPostRepositoryModelToEntity(v []Post) []*entities.Post {
// 	var posts []*entities.Post
// 	for _, post := range v {
// 		posts = append(posts, entities.NewPost(post.Id, post.Title, post.Body, post.UserId, post.Username, post.CreatedAt, post.UpdatedAt))
// 	}
// 	return posts
// }
