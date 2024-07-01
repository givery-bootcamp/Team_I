package repositories

import (
	"fmt"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type IntentionRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Intention struct {
	UserId int
	PostId int
	Status string
}

type IntentionRepositoryError struct {
	Message string
}

func (e *IntentionRepositoryError) Error() string {
	return e.Message
}

func (e *IntentionRepositoryError) IntentionError() string {
	return e.Message
}

func NewIntentionRepository(conn *gorm.DB) *IntentionRepository {
	return &IntentionRepository{
		Conn: conn,
	}
}

var ErrOnCreateIntention = &IntentionRepositoryError{Message: "error on create intention"}
var ErrOnDeleteIntention = &IntentionRepositoryError{Message: "error on delete intention"}

func (r *IntentionRepository) Create(userId int, postId int, status string) (*entities.IntentionForInsert, error) {
	intention := entities.IntentionForInsert{
		UserId: userId,
		PostId: postId,
		Status: status,
	}

	if err := r.Conn.Table("intentions").Select("user_id", "post_id", "status").Create(&intention).Error; err != nil {
		fmt.Println(err)
		return nil, ErrOnCreateIntention
	}

	return &entities.IntentionForInsert{
		UserId: userId,
		PostId: postId,
		Status: status,
	}, nil
}

func (r *IntentionRepository) Exists(userId int, postId int) (string, error) {
	var intention []Intention
	if err := r.Conn.Table("intentions").Where("post_id = ? and user_id = ?", postId, userId).
		Scan(&intention).Error; err != nil {
		return "", err
	}
	if len(intention) == 0 {
		return "", nil
	}
	if intention[0].Status != "attend" && intention[0].Status != "skip" {
		return "", fmt.Errorf("invalid status")
	}
	return intention[0].Status, nil
}

func (r *IntentionRepository) Delete(userId int, postId int) error {
	if err := r.Conn.Table("intentions").Where("post_id = ? and user_id = ?", postId, userId).
		Delete(&entities.IntentionForInsert{}).Error; err != nil {
		fmt.Println(err)
		return ErrOnDeleteIntention
	}
	return nil
}

func (r *IntentionRepository) GetUsersByPostIdAndStatus(postId int, status string) ([]*entities.User, error) {
	var users []User
	if err := r.Conn.Table("intentions").
		Select("users.id as id, users.name as name, users.password as password, users.created_at as created_at, users.updated_at as updated_at").
		Where("intentions.post_id = ? and intentions.status = ?", postId, status).
		Joins("JOIN users ON intentions.user_id = users.id").
		Scan(&users).Error; err != nil {
		return nil, err
	}
	result := []*entities.User{}
	for _, user := range users {
		result = append(result, entities.NewUser(user.Id, user.Name, user.Password, user.CreatedAt, user.UpdatedAt))
	}
	return result, nil
}
