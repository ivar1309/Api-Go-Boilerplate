-- name: CreateUser :one
INSERT INTO users (username, password_hash, role) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY username;

-- name: GetUserByUsername :one
SELECT id, username, password_hash, role 
FROM users 
WHERE username=$1;

-- name: UpdateUsersPassword :one
UPDATE users
SET password_hash=$2
WHERE username=$1
RETURNING *;

-- name: UpdateUsersRole :one
UPDATE users
SET role=$2
WHERE username=$1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username=$1;