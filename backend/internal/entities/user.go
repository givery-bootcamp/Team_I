package entities

type UserRepository interface {
	Create(username, password string) (*User, error)
	GetUserById(id int) (*User, error)
	GetUserByName(name string) (*User, error)
}

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(id int, name, password, createdAt, updatedAt string) *User {
	return &User{
		Id:        id,
		Name:      name,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
