
-- name: SessionsList :many 
SELECT 
    s.session_id,
    s.car_id,
    s.is_live,
    s.created_at,
    s.finished_at,
    COUNT(e.*) emergencies,
    COUNT(c.*) default_codes
FROM "sessions"  AS s
 
JOIN session_codes sc ON sc.session_id = s.session_id 
LEFT JOIN codes  e ON sc.code_id = e.code_id  AND e.is_emergency = TRUE
LEFT JOIN codes  c ON sc.code_id = c.code_id  AND c.is_emergency = FALSE
WHERE car_id = $1
GROUP BY 
s.session_id,
    s.car_id,
    s.is_live,
    s.created_at,
    s.finished_at;


-- name: SessionGetCodes :many 
SELECT c.* FROM session_codes sc JOIN codes c ON sc.code_id = c.code_id WHERE sc.session_id =  $1;

-- name: SessionCreate :one 
INSERT INTO
    "sessions" (
        car_id,
        is_live
    )
VALUES
    ($1 , $2) RETURNING *;




-- name: SessionAttachCodes :many
WITH code_ids AS (
  SELECT code_id FROM codes WHERE code_name = ANY(string_to_array(sqlc.arg('session_codes'), ','))
)
INSERT INTO session_codes (session_id, code_id)
SELECT sqlc.arg('session_id'), code_id FROM code_ids RETURNING *;



-- name: SessionAttachValues :many
INSERT INTO session_values (session_id, value_key, value_data)
SELECT sqlc.arg('session_id'), t.key, u.data
FROM unnest(string_to_array(sqlc.arg('session_keys'), ',')) WITH ORDINALITY AS t(key, n)
JOIN unnest(string_to_array(sqlc.arg('session_values'), ',')) WITH ORDINALITY AS u(data, n)
ON t.n = u.n
RETURNING *;
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