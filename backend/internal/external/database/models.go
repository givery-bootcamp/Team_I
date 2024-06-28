package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(40);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`

	Posts    []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Post struct {
	gorm.Model
	UserID int    `gorm:"not null"`
	Title  string `gorm:"type:varchar(100);not null"`
	Body   string `gorm:"type:text;not null"`
	Type   string `gorm:"type:varchar(20)"`
	Tag    string `gorm:"type:varchar(20)"`

	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Comment struct {
	gorm.Model
	UserID int    `gorm:"not null"`
	PostID int    `gorm:"not null"`
	Body   string `gorm:"type:text;not null"`
}
