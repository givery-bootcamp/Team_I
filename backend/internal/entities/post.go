package entities

type PostRepository interface {
	List() ([]*Post, error)
	GetPostById(id int) (*Post, error)
	Create(userId int, title, body string) (*PostForInsert, error)
}

type Post struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	UserId    int        `json:"user_id"`
	Username  string     `json:"username"`
	Comments  []*Comment `json:"comments"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type PostForInsert struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
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
