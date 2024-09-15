-- name: CreateUser :one
INSERT INTO user (id, name, password)
VALUES (?, ?, ?)
RETURNING *;