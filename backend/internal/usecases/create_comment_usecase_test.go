package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"myapp/internal/repositories"
)

func TestCreateCommentUsecaseExecute(t *testing.T) {
	userId := 1
	post_id := 1
	body := "test_body"

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	usecase := NewCreateCommentUsecase(r)
	result, err := usecase.Execute(userId, post_id, body)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userId, result.UserId)
	assert.Equal(t, post_id, result.PostId)
	assert.NotNil(t, result.CreatedAt)
	assert.NotNil(t, result.UpdatedAt)
}

func TestPostWithSpecifiedPostidNotFound(t *testing.T) {
	userId := 1
	post_id := 1000
	body := "test_body"

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	usecase := NewCreateCommentUsecase(r)
	result, err := usecase.Execute(userId, post_id, body)
	assert.Equal(t, err, &repositories.PostIdNotFound{})
	assert.Error(t, err)
	assert.Nil(t, result)
}
