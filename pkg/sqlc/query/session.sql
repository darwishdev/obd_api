
-- name: SessionsList :many 
SELECT * FROM "sessions" WHERE car_id = $1;

-- name: SessionCreate :one 
INSERT INTO
    "sessions" (
        car_id,
        is_live
    )
VALUES
    ($1 , $2) RETURNING *;




-- name: SessionAttachCode :one 
INSERT INTO
    "session_codes" (
        session_id,
        code_id
    )
VALUES
    ($1  , $2) RETURNING *;



-- name: SessionAttachValue :one 
INSERT INTO
    "session_values" (
        session_id,
        value_key,
        value_data
    )
VALUES
    ($1 , $2 , $3) RETURNING *;

-- name: SessionClose :one 
UPDATE
    "sessions"
SET
    finished_at = NOW()
WHERE
    session_id = $1 RETURNING *;



-- engine run time
--Max Rpm
--Max Speed
--Idling Fuel Consumtion
--Driving Fuel Consumtion fuel/per 100 km
--Rapid Acceleration Times (max speed)
--Rapid Decleration Times (framl_)
--Distance Travel
--Distance since codes cleare
-- Fuel Type Value (92/95)

--Instant Fuel Consumtion
--Intake pressure
--speed