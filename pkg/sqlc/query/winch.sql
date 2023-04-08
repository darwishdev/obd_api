

-- name: WinchCreate :one 
INSERT INTO
    winch (
        area_id,
        name,
        phone,
        driver_name,
        driver_phone
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;

-- name: WinchList :many
SELECT
    w.* , a.name area_name
FROM
    winch w JOIN areas a ON w.area_id = a.area_id
WHERE
    a.area_id = $1
    AND deleted_at IS NULL;