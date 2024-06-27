package repositories

import (
	"fmt"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

var ErrUserNotFound = fmt.Errorf("user not found")

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
	return entities.NewUser(user.Id, user.Name, user.Password, user.CreatedAt, user.UpdatedAt), nil
}

func (r *UserRepository) GetUserByName(name string) (*entities.User, error) {
	var result []User
	if err := r.Conn.Table("users").Select("id, name, password, created_at, updated_at").Where("name = ?", name).Scan(&result).Error; err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrUserNotFound
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("duplicate user")
	}
	// len(result) == 1
	user := result[0]
	return entities.NewUser(user.Id, user.Name, user.Password, user.CreatedAt, user.UpdatedAt), nil
}

func (r *UserRepository) Create(name, hashedPassword string) (*entities.User, error) {
	user := entities.User{
		Name:     name,
		Password: hashedPassword,
	}
	if err := r.Conn.Table("users").Select("name", "password").Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
