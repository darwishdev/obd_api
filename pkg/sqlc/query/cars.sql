-- name: CarBrandsList :many 
SELECT * FROM car_brands ;

-- name: CarBrandModelsList :many 
SELECT * FROM car_brand_models WHERE  car_brand_id = $1;


-- name: CarCreate :one 
INSERT INTO
    cars (
        car_brand_model_id,
        user_id,
        model_year
    )
VALUES
    ($1, $2, $3) RETURNING *;





-- name: CarUpdate :one 
UPDATE cars 
SET
    car_brand_model_id =coalesce(sqlc.narg('car_brand_model_id'), car_brand_model_id),
    model_year =coalesce(sqlc.narg('model_year'), model_year)
    WHERE car_id = sqlc.arg('car_id')
    RETURNING *;