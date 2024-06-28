package entities

import "time"

type PostRepository interface {
	List(page int, limit int, postType string) ([]*Post, error)
	GetPostById(id int) (*Post, error)
	DeletePost(id int) error
	Create(userId int, title, body, postType string) (*PostForInsert, error)
	Update(id int, title, body string) (*Post, error)
}

type Post struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Type      string     `json:"type"`
	UserId    int        `json:"user_id"`
	Username  string     `json:"username"`
	Comments  []*Comment `json:"comments,omitempty"`
	CreatedAt string     `json:"created_at" gorm:"-"`
	UpdatedAt string     `json:"updated_at" gorm:"-"`
}

type PostForInsert struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewPost(id int, title, body string, userId int, username, typ, createdAt, updatedAt string) *Post {
	return &Post{
		Id:        id,
		Title:     title,
		Body:      body,
		UserId:    userId,
		Username:  username,
		Type:      typ,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
