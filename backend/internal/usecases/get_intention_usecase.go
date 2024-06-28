package usecases

import (
	"myapp/internal/entities"
)

type GetIntentionUsecase struct {
	intentionRepository entities.IntentionRepository
	userRepository      entities.UserRepository
}

func NewGetIntentionUsecase(intensionRepo entities.IntentionRepository, userRepo entities.UserRepository) *GetIntentionUsecase {
	return &GetIntentionUsecase{
		intentionRepository: intensionRepo,
		userRepository:      userRepo,
	}
}

func (u *GetIntentionUsecase) Execute(postId int, status string) ([]*entities.User, error) {
	users, err := u.intentionRepository.GetUsersByPostIdAndStatus(postId, status)
	if err != nil {
		return nil, err
	}
	return users, nil
}
