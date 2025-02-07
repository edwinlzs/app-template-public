-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (id, alias, email) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET alias = $2
WHERE id = $1;
