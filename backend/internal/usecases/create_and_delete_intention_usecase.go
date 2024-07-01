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
	prevStatus, err := u.intentionRepository.Exists(userId, postId)
	if err != nil {
		return nil, WrapUsecaseError(err)
	}
	deleteFunc := func() (*entities.IntentionForInsert, error) {
		err := u.intentionRepository.Delete(userId, postId)
		if err != nil {
			return nil, WrapUsecaseError(err)
		}
		return nil, nil
	}
	createFunc := func() (*entities.IntentionForInsert, error) {
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
	if prevStatus == status {
		return deleteFunc()
	} else {
		if prevStatus == "" {
			return createFunc()
		} else {
			_, err := deleteFunc()
			if err != nil {
				return nil, WrapUsecaseError(err)
			}
			return createFunc()
		}
	}
}
