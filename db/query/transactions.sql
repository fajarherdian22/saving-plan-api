-- name: CreateTransaction :exec
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
);

-- name: GetTransaction :one
SELECT * FROM transactions_record
WHERE id = ?;