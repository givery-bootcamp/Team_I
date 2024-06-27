package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
)

type GetUserUsecase struct {
	repository entities.UserRepository
}

func NewGetUserUsecase(r *repositories.UserRepository) *GetUserUsecase {
	return &GetUserUsecase{
		repository: r,
	}
}

func (u *GetUserUsecase) Execute(userId int) (*entities.User, error) {

	// Call repository
	user, err := u.repository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
