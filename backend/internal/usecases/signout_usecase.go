package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PostSignoutUsecase struct {
	repository entities.UserRepository
}

func NewPostSignoutUsecase(r *repositories.UserRepository) *PostSignoutUsecase {
	return &PostSignoutUsecase{
		repository: r,
	}
}

func (u *PostSignoutUsecase) Execute() (string, error) {

	// トークンの発行（ヘッダー・ペイロード）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
