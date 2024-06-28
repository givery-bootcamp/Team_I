package usecases

import (
	"fmt"

	"github.com/pkg/errors"
)

type CommentNotFound struct{}

func (err *CommentNotFound) Error() string {
	return "Cannot edit comment. Comment not found."
}

type NoPermission struct{}

func (err *NoPermission) Error() string {
	return "cannot edit other user's comment"
}

const ErrPasswordIncorrectMessage = "password is incorrect"
const ErrUserNotFoundMessage = "user not found"
const ErrDuplicateUserMessage = "duplicate user"
const ErrOnCreateIntentionMessage = "error on create intention"
const ErrOnDeleteIntentionMessage = "error on delete intention"

var ErrPasswordIncorrect = fmt.Errorf(ErrPasswordIncorrectMessage)
var ErrUserNotFound = fmt.Errorf(ErrUserNotFoundMessage)

var KnownErrorsUser = map[string]int{
	ErrUserNotFoundMessage:  0,
	ErrDuplicateUserMessage: 0,
}
var KnownErrorsIntention = map[string]int{
	ErrOnCreateIntentionMessage: 0,
	ErrOnDeleteIntentionMessage: 0,
}

type ErrUserUsecase interface {
	UserError() string
}

type ErrIntentionUsecase interface {
	IntentionError() string
}

// ユースケースが知っているエラーかどうかを判定し、エラーをラップする
func WrapUsecaseError(err error) error {
	if err == nil {
		return nil
	}
	// type switchを使ってエラーは定義されたエラーかどうかを判定
	switch er := errors.Cause(err).(type) {
	case ErrUserUsecase:
		_, exists := KnownErrorsUser[er.UserError()]
		if exists {
			return err
		} else {
			return errors.Wrap(err, ErrUnknown.Error())
		}
	case ErrIntentionUsecase:
		_, exists := KnownErrorsIntention[er.IntentionError()]
		if exists {
			return err
		} else {
			return errors.Wrap(err, ErrUnknown.Error())
		}
	default:
		return errors.Wrap(err, ErrUnknown.Error())
	}
}
