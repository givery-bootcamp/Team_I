package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type PostSigninUsecase struct {
	repository entities.UserRepository
}

func NewPostSigninUsecase(r *repositories.UserRepository) *PostSigninUsecase {
	return &PostSigninUsecase{
		repository: r,
	}
}

func (u *PostSigninUsecase) Execute(username, password string) (*entities.User, error) {
	// Call repository
	user, err := u.repository.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, nil
	}
	return user, nil
}
