

-- name: WinchCreate :one 
INSERT INTO
    winch (
        area_id,
        name,
        phone,
        lat,
        long,
        driver_name,
        driver_phone
    )
VALUES
    ($1, $2, $3, $4, $5 , $6 , $7) RETURNING *;

-- name: WinchList :many
SELECT * FROM find_winch($1, $2);


-- name: WinchClean :exec
DELETE FROM winch;
