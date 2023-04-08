-- name: CenterCreate :one 
INSERT INTO
    centers (
        name,
        phone,
        location,
        address,
        area_id,
        lat,
        long
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7) RETURNING *;


-- name: CentersList :many
SELECT
    *
FROM
    centers
WHERE
    area_id = $1
    AND deleted_at IS NULL;