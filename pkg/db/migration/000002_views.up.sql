--views
CREATE VIEW car_info AS
SELECT
  c.user_id AS car_user_id,
  c.car_id,
  c.car_brand_model_id,
  cbm.car_brand_id,
  cb.name AS brand_name,
  cbm.name AS brand_model_name,
  c.model_year
FROM
  cars c
  JOIN car_brand_models cbm ON c.car_brand_model_id = cbm.car_brand_model_id
  JOIN car_brands cb ON cbm.car_brand_id = cb.car_brand_id;


--views
CREATE VIEW user_info AS
SELECT
  *
FROM
  users u
  LEFT JOIN car_info c ON u.user_id = c.car_user_id
WHERE
  deleted_at IS NULL;


--views
CREATE VIEW center_info AS
SELECT
  c.center_id,
  c.name,
  c.phone,
  c.image,
  c.address,
  c.area_id,
  c.lat,
  c.long,
  c.created_at,
  CAST(COALESCE(AVG(r.rate) , 0) AS float) avg_rate,
  COUNT(r.review_id) reviews_count,
  CAST(0 AS float4) distance
FROM
  centers c
  LEFT JOIN reviews r ON r.center_id = c.center_id
WHERE
  c.deleted_at IS NULL
GROUP BY
  c.center_id,
  c.name,
  c.phone,
  c.image,
  c.address,
  c.area_id,
  c.lat,
  c.long,
  c.created_at;

CREATE VIEW winch_info AS
SELECT
  w.winch_id,
  w.area_id,
  w.name,
  w.phone,
  w.driver_name,
  w.driver_phone,
  w.lat,
  w.long,
  w.created_at,
  CAST(0 AS float4) distance
FROM
  winch w
WHERE
  w.deleted_at IS NULL