

CREATE OR REPLACE FUNCTION user_create(
    name_arg TEXT,
    phone_arg TEXT,
    email_arg TEXT,
    password_arg TEXT,
    car_brand_model_id_arg INTEGER,
    model_year_arg INTEGER
)  RETURNS SETOF user_info AS $$ 
DECLARE created_user_id INTEGER;
BEGIN
INSERT INTO
    users (
        name,
        phone,
        email,
        password)
VALUES
    (   name_arg,
        phone_arg,
        email_arg,
        password_arg
    ) RETURNING user_id INTO created_user_id;

INSERT INTO
    cars (
        car_brand_model_id,
        model_year,
        user_id
    )
VALUES
    (
        car_brand_model_id_arg,
        model_year_arg,
        created_user_id
    );

    RETURN QUERY
    SELECT *
    FROM
        user_info u
    WHERE
        u.user_id = created_user_id;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION find_centers(
    in_lat float,
    in_long float
)
RETURNS  SETOF center_info AS $$
BEGIN
    RETURN QUERY
    WITH center_distances AS (
        SELECT c.center_id,
               c.name,
               c.phone,
               c.location,
               c.address,
               c.area_id,
               c.lat,
               c.long,
               c.created_at,
               c.avg_rate,
               c.reviews_count,
            CAST(point(c.lat,c.long) <-> point(in_lat,in_long) AS float4) AS distance
        FROM center_info c
    )
    SELECT *
    FROM center_distances
    ORDER BY distance;
END;
$$ LANGUAGE plpgsql;