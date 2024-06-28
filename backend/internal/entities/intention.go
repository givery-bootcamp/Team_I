package entities

type IntentionRepository interface {
	Create(userId int, postId int, state string) (*IntentionForInsert, error)
	Exists(userId int, postId int) (bool, error)
	Delete(userId int, postId int) error
	GetUsersByPostIdAndStatus(postId int, status string) ([]*User, error)
}

type Intention struct {
	UserId    int    `json:"user_id"`
	PostId    int    `json:"post_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IntentionForInsert struct {
	UserId   int    `json:"user_id"`
	PostId   int    `json:"post_id"`
	Status   string `json:"status"`
	UserName string `json:"user_name"`
}

func NewIntention(userId, postId int, state, createdAt, updatedAt string) *Intention {
	return &Intention{
		UserId:    userId,
		PostId:    postId,
		Status:    state,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
