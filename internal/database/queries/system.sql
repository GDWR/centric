-- name: InitialSetupComplete :one
SELECT COUNT(*) = 1 AS "InitialSetupComplete" 
FROM system;

-- name: InitialSetup :exec
INSERT INTO system (owner_id)
VALUES (?);