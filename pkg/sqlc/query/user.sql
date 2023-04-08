-- name: UserCreate :one
INSERT INTO
  users (name, phone, email, password)
VALUES
  ($1, $2, $3, $4) RETURNING *;

-- name: UserGet :one
SELECT
  *
FROM
  user_info u
WHERE
  u.user_id = $1;

-- name: UserGetByUsername :one
SELECT
  *
FROM
  user_info u
WHERE
  email = $1
  OR phone = $1;

-- name: UsersList :many
SELECT
  *
FROM
  users
WHERE
  deleted_at IS NULL
ORDER BY
  user_id
LIMIT
  $1 OFFSET $2;

-- name: UserUpdate :one
UPDATE
  users
SET
  name = coalesce(sqlc.narg('name'), name),
  phone = coalesce(sqlc.narg('phone'), phone),
  email = coalesce(sqlc.narg('email'), email),
  password = coalesce(sqlc.narg('password'), password),
  password_changed_at = coalesce(sqlc.narg('password'), password_changed_at)
WHERE
  user_id = sqlc.arg('user_id') RETURNING *;

-- name: UserDelete :one
UPDATE
  users
SET
  deleted_at = now()
WHERE
  user_id = $1 RETURNING *;