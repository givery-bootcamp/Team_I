package repositories

import (
	"fmt"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type User struct {
	Id        int
	Name      string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetUserById(id int) (*entities.User, error) {
	var user User
	if err := r.Conn.Table("users").Select("id, name, password, created_at, updated_at").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", user)
	return entities.NewUser(user.Id, user.Name, user.Password, user.CreatedAt, user.UpdatedAt), nil
}

func (r *UserRepository) GetUserByName(name string) (*entities.User, error) {
	var user User
	if err := r.Conn.Table("users").Select("id, name, password, created_at, updated_at").Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", user)
	return entities.NewUser(user.Id, user.Name, user.Password, user.CreatedAt, user.UpdatedAt), nil
}
