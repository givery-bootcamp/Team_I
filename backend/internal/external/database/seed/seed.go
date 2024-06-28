package seed

import (
	"gorm.io/gorm"
	"log"
	"myapp/internal/external/database"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64

	// Check if there are any users in the database
	if err := db.Model(&database.User{}).Count(&count).Error; err != nil {
		return err
	}

	// If there are already users, skip seeding.
	if count > 0 {
		return nil
	}

	users := []database.User{
		{Name: "taro", Password: "$2a$10$P9Zr9LES1Yv/n6k77pDy0OVwCRBeHRhHsFMQyU6GfkfpOXfHOPjgG"},
		{Name: "hanako", Password: "$2a$10$fyv7Ngey56irY9RkEG0MQuDzUyklKkWuqVV86gR3yTMxD6WeNZtZC"},
	}

	for i, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal("failed to seed users: ", result.Error)
		}

		posts := []database.Post{
			{UserID: int(user.ID), Title: "test1", Body: "質問1\n改行"},
			{UserID: int(user.ID), Title: "test2", Body: "質問2\n改行"},
			{UserID: int(user.ID), Title: "test3", Body: "質問1\n改行"},
		}

		if i == 0 {
			posts = append(posts, []database.Post{
				{UserID: int(user.ID), Title: "test4", Body: "質問2\n改行"},
				{UserID: int(user.ID), Title: "test7", Body: "質問1\n改行"},
				{UserID: int(user.ID), Title: "test9", Body: "質問1\n改行"},
				{UserID: int(user.ID), Title: "test10", Body: "質問2\n改行"},
			}...)
		} else {
			posts = append(posts, []database.Post{
				{UserID: int(user.ID), Title: "test5", Body: "質問1\n改行"},
				{UserID: int(user.ID), Title: "test8", Body: "質問2\n改行"},
			}...)
		}

		for _, post := range posts {
			result := db.Create(&post)
			if result.Error != nil {
				log.Fatal("failed to seed posts: ", result.Error)
			}

			comment := database.Comment{UserID: int(user.ID), PostID: int(post.ID), Body: "初コメ\n改行"}
			result = db.Create(&comment)
			if result.Error != nil {
				log.Fatal("failed to seed comments: ", result.Error)
			}
		}
	}
	return nil
}
