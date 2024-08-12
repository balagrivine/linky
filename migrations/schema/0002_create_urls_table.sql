-- +goose Up

CREATE TABLE urls (
	url_id INT PRIMARY KEY,
	original_url VARCHAR(255) NOT NULL,
	short_url VARCHAR(255) NOT NULL,
	user_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(user_id) REFERENCES users(user_id)
);

-- +goose Down
DROP TABLE urls;
