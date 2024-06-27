// このファイルはchihiroのコメント機能のテストファイルです。
// レイヤーごとのテストを書くのが面倒だったので、APIのテストを書いています。
// 今後はテスト用のデータベースサーバを建てる方針になったのでこちらのテストは使いませんが、
// sqlmockの使い方の資料として残しておきます。

package controllers_test

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"myapp/internal/controllers"
	"myapp/internal/entities"
	"myapp/internal/middleware"
	"time"

	// "myapp/internal/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func getTestUserJwtStr(id int, name, secret_key string) string {
	// トークンの発行（ヘッダー・ペイロード）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name": name,
		"Id":   id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func TestPostCommentController(t *testing.T) {
	// テストのためのデータ
	user_id := 1
	user_name := "taro"
	post_id := 2
	body := "test comment"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer db.Close()

	// GORMのモックデータベース接続を開く
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	// Mockを定義する
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `comments` (`user_id`,`post_id`,`body`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)").
		WithArgs(user_id, post_id, body, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// テスト用のHTTPリクエストを作成
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	middleware.SetupRoutes(r, gormDB)

	// テスト用のデータ
	postCommentInput := controllers.PostCommentInput{
		PostId: post_id,
		Body:   body,
	}
	jsonBytes, err := json.Marshal(postCommentInput)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/comments", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	tokenString := getTestUserJwtStr(user_id, user_name, middleware.SECRET_KEY)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: tokenString})

	r.ServeHTTP(w, req)

	var responseComment entities.Comment
	if err := json.Unmarshal(w.Body.Bytes(), &responseComment); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user_id, responseComment.UserId)
	assert.Equal(t, post_id, responseComment.PostId)
	assert.Equal(t, body, responseComment.Body)

}
