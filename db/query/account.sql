-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccountBalance :exec
UPDATE accounts 
SET balance = $2
WHERE id = $1;

-- name: UpdateAccountBalanceAndReturn :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;