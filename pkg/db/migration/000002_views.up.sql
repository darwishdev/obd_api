
--views
CREATE VIEW car_info AS
SELECT
    c.user_id AS user_id,
    cb.name AS brand_name,
    cbm.name AS brand_model_name,
    c.model_year,
    c.created_at
FROM
    cars c
    JOIN car_brand_models cbm ON c.car_brand_model_id = cbm.car_brand_model_id
    JOIN car_brands cb ON cbm.car_brand_id = cb.car_brand_id;

--views
CREATE VIEW user_info AS
SELECT
    u.*,
    c.brand_name,
    c.brand_model_name,
    c.model_year,
    c.created_at car_created_at
FROM
    users u
    LEFT JOIN car_info c ON u.user_id = c.user_id
WHERE
    deleted_at IS NULL;
