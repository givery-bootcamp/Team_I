package usecases

import (
	"fmt"
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
	_, err := u.repository.GetUserByName(name)
	// 既にユーザが存在する場合はエラーを返す
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	// ユーザが存在しない場合以外のエラーはそのまま返す
	if err != ErrUserNotFound {
		return nil, err
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

	return u.repository.Create(name, hashedPassword)
}
