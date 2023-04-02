-- name: UserCreate :one
INSERT INTO users (
  name,
  phone,
  email,
  password
) VALUES (
  $1, $2,$3, $4
) RETURNING *;


-- name: UserGet :one
SELECT * FROM users
WHERE user_id = $1 AND deleted_at IS NULL LIMIT 1;



-- name: UserGetByUsername :one
SELECT * FROM users
WHERE (email = $1 OR phone = $1) AND deleted_at IS NULL;


-- name: UsersList :many
SELECT * FROM users WHERE deleted_at IS NULL 
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UserUpdate :one
UPDATE users
SET "name" = $2 , "phone" = $3 ,  "email" = $4 , "password" = $5
WHERE user_id = $1
RETURNING *;


-- name: UserDelete :one
UPDATE users
SET deleted_at = now()
WHERE user_id = $1
RETURNING *;