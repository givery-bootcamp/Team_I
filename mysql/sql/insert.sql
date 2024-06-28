func
seedDatabase(db *gorm.DB) {
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
			{UserID: user.ID, Title: "test1", Body: "質問1\n改行"},
			{UserID: user.ID, Title: "test2", Body: "質問2\n改行"},
			{UserID: user.ID, Title: "test3", Body: "質問1\n改行"},
		}

		if i == 0 {
			posts = append(posts, []database.Post{
				{UserID: user.ID, Title: "test4", Body: "質問2\n改行"},
				{UserID: user.ID, Title: "test7", Body: "質問1\n改行"},
				{UserID: user.ID, Title: "test9", Body: "質問1\n改行"},
				{UserID: