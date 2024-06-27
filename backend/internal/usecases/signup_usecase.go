package usecases

import (
	"fmt"
	"log"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"regexp"

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
var ErrNotASCII = fmt.Errorf("password must be ASCII")
var ErrNotAlphaNumeric = fmt.Errorf("name must be alphanumeric")
var ErrUnexpected = fmt.Errorf("unexpected error")

func isAlphaNumeric(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return re.MatchString(s)
}

func isASCII(s string) bool {
	re := regexp.MustCompile("^[\x20-\x7E]*$")
	return re.MatchString(s)
}

func (u *PostSignupUsecase) Execute(name, password string) (*entities.User, error) {
	// ユーザ名が英数字のみから成るかチェック
	if !isAlphaNumeric(name) {
		return nil, ErrNotAlphaNumeric
	}
	// ユーザ名が既に存在するかチェック
	userExists, err := u.repository.UserExists(name)
	// 予期してないエラーの場合Unexpectedを返す
	if err != nil {
		log.Println(err)
		return nil, ErrUnexpected
	}
	// 既にユーザが存在する場合はエラーを返す
	if userExists {
		return nil, ErrUserAlreadyExists
	}

	// パスワードがASCII範囲の英数記号のみから成るかチェック
	if !isASCII(password) {
		return nil, ErrNotASCII
	}

	// 弱いパスワードを弾く
	if len(password) < 8 {
		return nil, ErrPasswordTooShort
	}

	passwordByte := []byte(password)
	hashedPasswordByte, _ := bcrypt.GenerateFromPassword(passwordByte, 10)
	hashedPassword := string(hashedPasswordByte)

	user, err := u.repository.Create(name, hashedPassword)
	if err != nil {
		log.Println(err)
		return nil, ErrUnexpected
	}
	return user, nil
}
