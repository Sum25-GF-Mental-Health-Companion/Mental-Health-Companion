-- name: CreateUser :one
INSERT INTO users (email, hashed_password, student_status)
VALUES ($1, $2, $3)
RETURNING id, email, hashed_password, student_status, created_at;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
