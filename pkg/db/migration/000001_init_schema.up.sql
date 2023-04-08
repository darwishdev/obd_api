CREATE TABLE "users" (
    "user_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "phone" character varying(200) UNIQUE NOT NULL,
    "email" character varying(200) UNIQUE NOT NULL,
    "password" character varying(200) NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00'),
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "deleted_at" timestamptz
);

CREATE TABLE "car_brands" (
    "car_brand_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL
);

CREATE TABLE "car_brand_models" (
    "car_brand_model_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "start_year" date NOT NULL,
    "end_year" date,
    "car_brand_id" bigint NOT NULL
);

CREATE TABLE "cars" (
    "car_id" bigserial PRIMARY KEY,
    "car_brand_model_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "model_year" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "deleted_at" timestamptz
);

CREATE TABLE "areas" (
    "area_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL
);

CREATE TABLE "centers" (
    "center_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "phone" character varying(200) NOT NULL,
    "location" character varying(200) NOT NULL,
    "address" character varying(200) NOT NULL,
    "area_id" bigint NOT NULL,
    "lat" character varying(200) NOT NULL,
    "long" character varying(200) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "deleted_at" timestamptz
);

CREATE TABLE "reviews" (
    "review_id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "center_id" bigint NOT NULL,
    "review" character varying(200) NOT NULL,
    "rate" smallint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "sessions" (
    "session_id" bigserial PRIMARY KEY,
    "car_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "finished_at" timestamptz
);

CREATE TABLE "session_results" (
    "session_result_id" bigserial PRIMARY KEY,
    "session_id" bigint NOT NULL,
    "code_name" character varying(200) NOT NULL
);

CREATE TABLE "codes" (
    "code_id" bigserial PRIMARY KEY,
    "session_id" bigint NOT NULL,
    "car_brand_model_id" bigint NOT NULL,
    "code_name" character varying(200) NOT NULL,
    "description" character varying(200) NOT NULL,
    "isEmergency" BOOLEAN
);

CREATE TABLE "winch" (
    "winch_id" bigserial PRIMARY KEY,
    "area_id" bigint NOT NULL,
    "name" character varying(200) NOT NULL,
    "phone" character varying(200) NOT NULL,
    "driver_name" character varying(200) NOT NULL,
    "driver_phone" character varying(200) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "deleted_at" timestamptz
);


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
WHERE deleted_at IS NULL;

-- constraints
ALTER TABLE
    "car_brand_models"
ADD
    FOREIGN KEY ("car_brand_id") REFERENCES "car_brands" ("car_brand_id");

ALTER TABLE
    "cars"
ADD
    FOREIGN KEY ("car_brand_model_id") REFERENCES "car_brand_models" ("car_brand_model_id");

ALTER TABLE
    "cars"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE
    "reviews"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE
    "reviews"
ADD
    FOREIGN KEY ("center_id") REFERENCES "centers" ("center_id");

ALTER TABLE
    "centers"
ADD
    FOREIGN KEY ("area_id") REFERENCES "areas" ("area_id");

ALTER TABLE
    "winch"
ADD
    FOREIGN KEY ("area_id") REFERENCES "areas" ("area_id");

--seeds

INSERT INTO
    areas (name)
VALUES
    ('Al Sheikh zayed'),
    ('Al mohandseen');
INSERT INTO
    car_brands (name)
VALUES
    ('Toyota'),
    ('Mercides'),
    ('Hyundai');

INSERT INTO
    car_brand_models (name, start_year, car_brand_id)
VALUES
    ('Corolla', '2010-01-01', 1),
    ('Elantra', '2010-01-01', 2);


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
    (
        'GB Auto - Hyundai - El Sheikh Zayed',
        '01207661993',
        '3X64+5W قسم الشيخ زايد',
        'خلف الهاييبر وان بجوار امريكانا بلازا - الحى الاول شارع الروضة امام كليه, الهندسة, Giza, الجيزة',
        1,
        '30.03735564',
        '31.01083698'
    ),
    (
        'مركز صيانة هيونداى Hyundai Service Center',
        '01006888826',
        '3X64+5W قسم الشيخ زايد',
        'بوابه, 8 بيفرلى هيلز، الجيزة 12588',
        1,
        '30.06500699',
        '30.95729685'
    ),
    (
        'GB Auto - Hyundai - Abo Rawash - Showroom, Service Center & spare parts',
        '0235366010',
        '329P+4X قسم الشيخ زايد',
        'ابورواش 28 طريق القاهرة الاسكندرية الصحراوى - المنطقة الصناعية الثانية - طريق ال, Abu Rawash, الجيزة',
        1,
        '30.067845596254937',
        '31.03740789'
    ),
    (
        'Wahdan Auto Group',
        '0238204410',
        'WV7W+J2 قسم ثان 6 أكتوبر',
        'المنطقة الصناعية الثالثة، السادس من اكتوبر، الجيزة، 3221732',
        1,
        '29.914065410175006',
        '30.89492614'
    ),
    (
        'مركز عربيتي',
        '01157555959',
        '2XV7+CR قسم الشيخ زايد',
        'رقم 45، قسم الشيخ زايد، الجيزة 3242803',
        1,
        '30.043544240746243',
        '30.96450447'
    );



INSERT INTO "winch" ("area_id", "name", "phone", "driver_name", "driver_phone")
VALUES
    (1, 'Winch 1', '1234567890', 'Driver 1', '1234567890'),
    (1, 'Winch 2', '1234567890', 'Driver 2', '1234567890'),
    (1, 'Winch 3', '1234567890', 'Driver 3', '1234567890'),
    (2, 'Winch 4', '1234567890', 'Driver 4', '1234567890'),
    (2, 'Winch 5', '1234567890', 'Driver 5', '1234567890');

