USE training;

ALTER TABLE users ADD CONSTRAINT unique_name UNIQUE (`name`);