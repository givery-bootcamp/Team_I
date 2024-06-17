package entities

type PostRepository interface {
	List() ([]*Post, error)
	GetPostById(id int) (*Post, error)
}

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewPost(id int, title, body string, userId int, username, createdAt, updatedAt string) *Post {
	return &Post{
		Id:        id,
		Title:     title,
		Body:      body,
		UserId:    userId,
		Username:  username,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
