-- +goose Up
CREATE TABLE users (
	user_id SERIAL NOT NULL,
	email VARCHAR(255) NOT NULL,
	PRIMARY KEY(user_id)
);

-- +goose Down
