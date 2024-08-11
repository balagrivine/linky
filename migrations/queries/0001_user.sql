-- name: CreateUser :one
INSERT INTO users("email") VALUES($1)
RETURNING email;

-- name: CreateURL :one
INSERT INTO urls("url_id", "original_url", "short_url", "user_id", "created_at") VALUES($1, $2, $3, $4, $5)
RETURNING url_id, original_url, short_url, user_id, created_at;

-- name: GetOriginalURL :one
SELECT original_url FROM urls WHERE original_url = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
