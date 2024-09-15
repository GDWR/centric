-- name: GetEnvironmentByID :one
SELECT *
FROM environment
WHERE id = ?;

-- name: GetEnvironmentByName :one
SELECT *
FROM environment
WHERE name = ?;

-- name: CreateEnvironment :one
INSERT INTO environment (id, name, uri, icon)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetEnvironments :many
SELECT *
FROM environment;