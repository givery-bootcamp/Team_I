package usecases

import (
	"myapp/internal/entities"
)

type CreateIntentionUsecase struct {
	intentionRepository entities.IntentionRepository
	userRepository      entities.UserRepository
}

func NewCreateIntentionUsecase(intensionRepo entities.IntentionRepository, userRepo entities.UserRepository) *CreateIntentionUsecase {
	return &CreateIntentionUsecase{
		intentionRepository: intensionRepo,
		userRepository:      userRepo,
	}
}

func (u *CreateIntentionUsecase) Execute(userId int, postId int, status string) (*entities.IntentionForInsert, error) {
	exists, err := u.intentionRepository.Exists(userId, postId)
	if err != nil {
		return nil, WrapUsecaseError(err)
	}
	if exists {
		err := u.intentionRepository.Delete(userId, postId)
		if err != nil {
			return nil, WrapUsecaseError(err)
		}
		return nil, nil
	} else {
		intention, err := u.intentionRepository.Create(userId, postId, status)
		if err != nil {
			return nil, WrapUsecaseError(err)
		}
		user, err := u.userRepository.GetUserById(userId)
		if err != nil {
			return nil, WrapUsecaseError(err)
		}
		intention.UserName = user.Name
		return intention, nil
	}
}
