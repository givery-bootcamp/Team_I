package entities

type CommentRepository interface {
	GetListByPostId() ([]*Comment, error)
	// CreateComment(comment *Comment) error
	// UpdateComment(comment *Comment) error
	// DeleteComment(id int) error
}

type Comment struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	PostId    int    `json:"post_id"`
	UserName  string `json:"user_name"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewComment(id, user_id, post_id int, user_name, body, createdAt, updatedAt string) *Comment {
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