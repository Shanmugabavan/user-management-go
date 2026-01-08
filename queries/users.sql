-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    phone,
    age,
    status
)
VALUES ( $1, $2, $3, $4, $5, $6)
    RETURNING user_id, email, status;