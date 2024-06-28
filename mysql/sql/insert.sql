INSERT INTO users (name, password) VALUES ('taro', '$2a$10$P9Zr9LES1Yv/n6k77pDy0OVwCRBeHRhHsFMQyU6GfkfpOXfHOPjgG');
INSERT INTO users (name, password) VALUES ('hanako', '$2a$10$fyv7Ngey56irY9RkEG0MQuDzUyklKkWuqVV86gR3yTMxD6WeNZtZC');

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'test2', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test3', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test4', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'test5', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test6', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test7', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'test8', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test9', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'test10', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest1', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'atest2', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest3', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest4', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'atest5', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest6', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest7', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (2, 'atest8', '質問2\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest9', '質問1\n改行');
INSERT INTO posts (user_id, title, body)
VALUES (1, 'atest10', '質問2\n改行');

INSERT INTO comments (user_id, post_id, body) VALUES (2, 1, '初コメ\n改行');
INSERT INTO comments (user_id, post_id, body) VALUES (2, 1, '2コメ\n改行');
