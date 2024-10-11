-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email,hashed_password)
VALUES (
    gen_random_uuid (),
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    $1,
    $2
)
RETURNING *;

-- name: UpdateUsers :exec
UPDATE users
SET email = $1, hashed_password = $2
WHERE id = $3;

-- name: GetUser :one
SELECT * FROM users WHERE $1 = users.email;
-- name: DeleteUsers :exec
delete  from users;