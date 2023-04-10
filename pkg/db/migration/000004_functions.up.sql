

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