-- +goose Up
CREATE TABLE users (
	user_id INT PRIMARY KEY,
	email VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;
