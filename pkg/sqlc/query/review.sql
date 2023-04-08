-- name: ReviewCreate :one 
INSERT INTO
    reviews (user_id, center_id, review, rate)
VALUES
    ($1,$2,$3,$4) RETURNING *;

-- name: ReviewsList :many
SELECT
    u.name user_name , c.name center_name , r.review , r.rate , r.created_at
FROM
    reviews r
    JOIN users u ON r.user_id = u.user_id
    JOIN centers c ON r.center_id = c.center_id
WHERE
    r.user_id = coalesce(sqlc.narg('user_id'), r.user_id)
    AND r.center_id = coalesce(sqlc.narg('center_id'), r.center_id);