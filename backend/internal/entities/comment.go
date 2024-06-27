package entities

import (
	"time"
)

type CommentRepository interface {
	GetById(comment_id int) (*Comment, error)
	GetListByPostId(post_id int) ([]*Comment, error)
	Create(comment *Comment) (*Comment, error)
	Update(comment *Comment) (*Comment, error)
	Delete(comment_id int) error
}

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PostId    int       `json:"post_id"`
	UserName  string    `json:"user_name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewComment(id, user_id, post_id int, user_name, body string, createdAt, updatedAt time.Time) *Comment {
	return &Comment{
		Id:        id,
		UserId:    user_id,
		PostId:    post_id,
		UserName:  user_name,
		Body:      body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
