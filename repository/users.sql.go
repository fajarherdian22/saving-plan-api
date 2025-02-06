// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package repository

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    created_at
) VALUES (
    ?, ?, ?
)
`

type CreateUserParams struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.ID, arg.Email, arg.CreatedAt)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, created_at FROM users
WHERE email = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.CreatedAt)
	return i, err
}
