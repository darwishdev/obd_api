-- name: AreaCreate :one 
INSERT INTO areas (name) VALUES ($1) RETURNING *;



-- name: AreasList :many
SELECT * FROM areas;



-- name: AreaDelete :exec
DELETE  FROM areas WHERE area_id = $1;
