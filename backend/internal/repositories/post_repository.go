package repositories

import (
	"fmt"
	"time"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Post struct {
	Id        int
	Title     string
	Body      string
	UserId    int
	Username  string
	CreatedAt string
	UpdatedAt string
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) List() ([]*entities.Post, error) {
	var posts []Post
	if err := r.Conn.Table("posts").Select("posts.id, users.name as username, posts.user_id, posts.title, posts.body, posts.created_at, posts.updated_at").Where("posts.deleted_at IS NULL").Joins("JOIN users ON posts.user_id = users.id").Order("posts.id DESC").Scan(&posts).Error; err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", posts)
	return convertPostRepositoryModelToEntity(posts), nil
}

func (r *PostRepository) GetPostById(id int) (*entities.Post, error) {
	var post Post
	if err := r.Conn.Table("posts").Select("posts.id, users.name as username, posts.user_id, posts.title, posts.body, posts.created_at, posts.updated_at").Joins("JOIN users ON posts.user_id = users.id").Where("posts.id = ? AND posts.deleted_at IS NULL", id).First(&post).Error; err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", post)
	return entities.NewPost(post.Id, post.Title, post.Body, post.UserId, post.Username, post.CreatedAt, post.UpdatedAt), nil
}

func (r *PostRepository) DeletePost(id int) error {
	t := time.Now()
	//　レコードを論理削除する
	if err := r.Conn.Table("posts").Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", t.Format(time.DateTime)).Error; err != nil {
		return err
	}
	return nil
}

func convertPostRepositoryModelToEntity(v []Post) []*entities.Post {
	var posts []*entities.Post
	for _, post := range v {
		posts = append(posts, entities.NewPost(post.Id, post.Title, post.Body, post.UserId, post.Username, post.CreatedAt, post.UpdatedAt))
	}
	return posts
}
