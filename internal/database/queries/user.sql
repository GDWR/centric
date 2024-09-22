-- name: CreateUser :one
INSERT INTO user (id, name, password)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByName :one
SELECT id, name, password
FROM user
WHERE name = ?;