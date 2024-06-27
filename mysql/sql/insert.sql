INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');

INSERT INTO users (name, password) VALUES ('taro', '$2a$10$P9Zr9LES1Yv/n6k77pDy0OVwCRBeHRhHsFMQyU6GfkfpOXfHOPjgG');
INSERT INTO users (name, password) VALUES ('hanako', '$2a$10$fyv7Ngey56irY9RkEG0MQuDzUyklKkWuqVV86gR3yTMxD6WeNZtZC');

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test2', '質問2\n改行');

INSERT INTO comments (user_id, post_id, body) VALUES (2, 1, '初コメ\n改行');
INSERT INTO comments (user_id, post_id, body) VALUES (2, 1, '2コメ\n改行');
