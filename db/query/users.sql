-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    created_at
) VALUES (
    ?, ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE email = ? LIMIT 1;