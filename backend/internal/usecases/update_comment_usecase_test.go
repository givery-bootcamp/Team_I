package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"myapp/internal/entities"
	"myapp/internal/repositories"
)

func TestUpdateCommentUsecaseExecute(t *testing.T) {
	user_id := 2
	body := "test_body"

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	// ダミーデータ作成
	comment := &entities.Comment{
		UserId: user_id,
		PostId: 1,
		Body:   "test_body",
	}
	_, err := r.Create(comment)
	assert.NoError(t, err)
	comment_id := comment.Id

	usecase := NewUpdateCommentUsecase(r)

	result, err := usecase.Execute(user_id, comment_id, body)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, comment_id, result.Id)
	assert.Equal(t, body, result.Body)
	assert.NotNil(t, result.CreatedAt)
	assert.NotNil(t, result.UpdatedAt)
}

func TestUpdateCommentNotFound(t *testing.T) {
	user_id := 1
	comment_id := 1000
	body := "test_body"

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	usecase := NewUpdateCommentUsecase(r)
	result, err := usecase.Execute(user_id, comment_id, body)
	assert.Error(t, err)
	assert.Equal(t, err, &CommentNotFound{})
	assert.Nil(t, result)
}

func TestUpdateNoPermission(t *testing.T) {
	create_user_id := 1
	update_user_id := 2
	body := "test_body"

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	// ダミーデータ作成
	comment := &entities.Comment{
		UserId: create_user_id,
		PostId: 1,
		Body:   "test_body",
	}
	_, err := r.Create(comment)
	assert.NoError(t, err)
	comment_id := comment.Id

	usecase := NewUpdateCommentUsecase(r)
	result, err := usecase.Execute(update_user_id, comment_id, body)
	assert.Error(t, err)
	assert.Equal(t, err, &NoPermission{})
	assert.Nil(t, result)
}
