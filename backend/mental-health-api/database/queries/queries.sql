-- name: CreateUser :one
INSERT INTO users (email, hashed_password, student_status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: StartSession :one
INSERT INTO sessions (user_id) VALUES ($1)
RETURNING *;

-- name: EndSession :exec
UPDATE sessions SET ended_at = now() WHERE id = $1;

-- name: AddMessage :one
INSERT INTO messages (session_id, sender, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: SaveSummary :exec
INSERT INTO summaries (session_id, full_summary, compressed_summary)
VALUES ($1, $2, $3)
ON CONFLICT (session_id) DO UPDATE
SET full_summary = EXCLUDED.full_summary,
    compressed_summary = EXCLUDED.compressed_summary;

-- name: GetSessionsByUser :many
SELECT * FROM sessions WHERE user_id = $1 ORDER BY started_at DESC;
