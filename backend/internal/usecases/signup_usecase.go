package usecases

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type PostSignupUsecase struct {
	repository entities.UserRepository
}

func NewPostSignupUsecase(r *repositories.UserRepository) *PostSignupUsecase {
	return &PostSignupUsecase{
		repository: r,
	}
}

var ErrUserAlreadyExists = fmt.Errorf("user name already exists")
var ErrPasswordTooShort = fmt.Errorf("password is too short")

func (u *PostSignupUsecase) Execute(name, password string) (*entities.User, error) {
	// ユーザ名が既に存在するかチェック
	_, err := u.repository.GetUserByName(name)
	// 既にユーザが存在する場合はエラーを返す
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	// ユーザが存在しない場合以外のエラーはそのまま返す
	if err != ErrUserNotFound {
		return nil, err
	}

	// 弱いパスワードを弾く
	if len(password) < 8 {
		return nil, ErrPasswordTooShort
	}

	passwordByte := []byte(password)
	hashedPasswordByte, _ := bcrypt.GenerateFromPassword(passwordByte, 10)
	hashedPassword := string(hashedPasswordByte)

	return u.repository.Create(name, hashedPassword)
}
