-- name: UserCreate :one
SELECT * FROM user_create(
sqlc.arg('name_arg'),
sqlc.arg('phone_arg'),
sqlc.arg('email_arg'),
sqlc.arg('password_arg'),
sqlc.arg('car_brand_model_id_arg'),
sqlc.arg('model_year_arg'));

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
  password_changed_at = CASE
                          WHEN sqlc.narg('password') IS NULL THEN password_changed_at
                          ELSE now()
                        END
WHERE
  user_id = sqlc.arg('user_id') RETURNING *;

-- name: UserDelete :one
UPDATE
  users
SET
  deleted_at = now()
WHERE
  user_id = $1 RETURNING *;