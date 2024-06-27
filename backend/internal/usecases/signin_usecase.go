package usecases

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const errPasswordIncorrectMessage = "password is incorrect"
const errUserNotFoundMessage = "user not found"
const errUnknownMessage = "unknown error"

var ErrPasswordIncorrect = fmt.Errorf(errPasswordIncorrectMessage)
var ErrUserNotFound = fmt.Errorf(errUserNotFoundMessage)
var ErrUnknown = fmt.Errorf(errUnknownMessage)

type ErrSigninUsecase interface {
	SigninError() string
}

func WrapSigninUsecaseError(err error) error {
	if err == nil {
		return nil
	}
	er, ok := errors.Cause(err).(ErrSigninUsecase)
	if !ok {
		return errors.Wrap(err, ErrUnknown.Error())
	}
	switch er.SigninError() {
	case errPasswordIncorrectMessage:
		return errors.Wrap(err, ErrPasswordIncorrect.Error())
	case errUserNotFoundMessage:
		return errors.Wrap(err, ErrUserNotFound.Error())
	default:
		return errors.Wrap(err, ErrUnknown.Error())
	}
}

type PostSigninUsecase struct {
	repository entities.UserRepository
}

func NewPostSigninUsecase(r entities.UserRepository) *PostSigninUsecase {
	return &PostSigninUsecase{
		repository: r,
	}
}

func (u *PostSigninUsecase) Execute(username, password string) (*entities.User, string, error) {

	// Call repository
	user, err := u.repository.GetUserByName(username)
	if err != nil {
		return nil, "", WrapSigninUsecaseError(err)
	}
	passwordByte := []byte(password)
	storedPasswordByte := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(storedPasswordByte, passwordByte)

	if err != nil {
		return nil, "", ErrPasswordIncorrect
	}

	// トークンの発行（ヘッダー・ペイロード）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name": user.Name,
		"Id":   user.Id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
