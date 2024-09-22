-- name: InitialSetupComplete :one
SELECT COUNT(*) = 1 AS "InitialSetupComplete" 
FROM system;

-- name: InitialSetup :exec
INSERT INTO system (owner_id, jwt_secret)
VALUES (?, ?);

-- name: GetJwtSecret :one
SELECT jwt_secret
FROM system;