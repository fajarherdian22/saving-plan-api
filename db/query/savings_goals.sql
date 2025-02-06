-- name: CreateSavingPlan :exec
INSERT INTO saving_goals (
    id,
    user_id,
    name,
    target_amount,
    current_amount,
    target_time,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, NOW()
);

-- name: GetSavingPlan :one
SELECT * FROM saving_goals
WHERE user_id = ? AND id = ? LIMIT 1;