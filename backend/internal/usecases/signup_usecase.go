package usecases

import (
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

func (u *PostSignupUsecase) Execute(name, password string) (*entities.User, error) {

	passwordByte := []byte(password)
	hashedPasswordByte, _ := bcrypt.GenerateFromPassword(passwordByte, 10)
	hashedPassword := string(hashedPasswordByte)

	return u.repository.Create(name, hashedPassword)
}
