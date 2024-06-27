package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"myapp/internal/entities"
	"myapp/internal/repositories"
)

func TestDeleteCommentUsecaseExecute(t *testing.T) {
	user_id := 2

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
	usecase := NewDeleteCommentUsecase(r)
	err = usecase.Execute(user_id, comment_id)

	assert.NoError(t, err)
}

func TestDeleteCommentNotFound(t *testing.T) {
	user_id := 1
	comment_id := 1000

	db := SetupDB()

	r := repositories.NewCommentRepository(db)
	usecase := NewDeleteCommentUsecase(r)
	err := usecase.Execute(user_id, comment_id)
	assert.Error(t, err)
	assert.Equal(t, err, &CommentNotFound{})
}

func TestDeleteNoPermission(t *testing.T) {
	create_user_id := 2
	delete_user_id := 1

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
	usecase := NewDeleteCommentUsecase(r)

	err = usecase.Execute(delete_user_id, comment_id)
	assert.Error(t, err)
	assert.Equal(t, err, &NoPermission{})
}
