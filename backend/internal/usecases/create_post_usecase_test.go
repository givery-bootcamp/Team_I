package usecases

import (
	"testing"

	"fmt"
	"myapp/internal/config"
	"myapp/internal/repositories"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Setup
// !!! You have to call this function after config setup
func SetupDB() *gorm.DB {
	host := "db-test"
	port := config.DBPort
	dbname := config.DBName
	dbUsername := config.DBUsername
	dbPassword := config.DBPassword
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}

func TestExecute(t *testing.T) {
	userId := 1
	title := "test_title"
	body := "test_body"
	postType := ""

	db := SetupDB()

	r := repositories.NewPostRepository(db)
	usecase := NewCreatePostUsecase(r)
	result, err := usecase.Execute(userId, title, body, postType)
	if err != nil {
		t.Fail()
	}

	if result.UserId != userId || result.Title != title || result.Body != body {
		t.Fail()
	}
}
