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
const ErrUserNotFoundMessage = "user not found"

var ErrPasswordIncorrect = fmt.Errorf(errPasswordIncorrectMessage)
var ErrUserNotFound = fmt.Errorf(ErrUserNotFoundMessage)

type ErrSigninUsecase interface {
	SigninError() string
}

// ユースケースが知っているエラーかどうかを判定し、エラーをラップする
func WrapSigninUsecaseError(err error) error {
	if err == nil {
		return nil
	}
	// インターフェースを使ってエラーを判定
	er, ok := errors.Cause(err).(ErrSigninUsecase)
	// インターフェースを実装してなければ、知らないエラーとして扱う
	if !ok {
		return errors.Wrap(err, ErrUnknown.Error())
	}
	// 定義されたエラーかどうかを判定
	switch er.SigninError() {
	case ErrUserNotFoundMessage:
		return errors.Wrap(err, ErrUserNotFound.Error())
	// ここは本来実行されない。リポジトリ側の実装が不適切だった場合に実行される
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
