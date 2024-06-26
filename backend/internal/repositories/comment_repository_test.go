package repositories

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// GORMのモックデータベース接続を開く
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	// CommentRepositoryのインスタンスを作成
	repo := NewCommentRepository(gormDB)

	// テストデータ
	userID := 1
	postID := 1
	body := "This is a comment"

	// モックの期待される動作を設定
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `comments`").WithArgs(userID, postID, body, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 関数を実行
	result, err := repo.Create(userID, postID, body)

	// アサーション
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, userID, result.UserId)
	assert.Equal(t, postID, result.PostId)
	assert.NotNil(t, result.CreatedAt)
	assert.NotNil(t, result.UpdatedAt)
	assert.Equal(t, body, result.Body)

	// すべてのモックの期待が満たされていることを確認
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
