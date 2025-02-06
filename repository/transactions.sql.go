// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transactions.sql

package repository

import (
	"context"
	"database/sql"
)

const createTransaction = `-- name: CreateTransaction :exec
INSERT INTO transactions_record (
    id,
    user_id,
    goal_id,
    amount,
    type,
    category,
    name,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, NOW()
)
`

type CreateTransactionParams struct {
	ID       string                 `json:"id"`
	UserID   string                 `json:"user_id"`
	GoalID   sql.NullString         `json:"goal_id"`
	Amount   float64                `json:"amount"`
	Type     TransactionsRecordType `json:"type"`
	Category string                 `json:"category"`
	Name     string                 `json:"name"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, createTransaction,
		arg.ID,
		arg.UserID,
		arg.GoalID,
		arg.Amount,
		arg.Type,
		arg.Category,
		arg.Name,
	)
	return err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, user_id, goal_id, amount, type, category, name, created_at FROM transactions_record
WHERE id = ?
`

func (q *Queries) GetTransaction(ctx context.Context, id string) (TransactionsRecord, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i TransactionsRecord
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.GoalID,
		&i.Amount,
		&i.Type,
		&i.Category,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}
