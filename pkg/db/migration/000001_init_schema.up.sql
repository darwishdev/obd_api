CREATE TABLE "areas" (
    "area_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL
);

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
    "years" TEXT NOT NULL,
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

CREATE TABLE "centers" (
    "center_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "phone" character varying(200) NOT NULL,
    "location" character varying(200) NOT NULL,
    "address" character varying(200) NOT NULL,
    "area_id" bigint NOT NULL,
    "lat" REAL NOT NULL,
    "long" REAL NOT NULL,
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
    "is_live" BOOLEAN   NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "finished_at" timestamptz
);
COMMENT ON COLUMN sessions.is_live IS 'indicates if session loads live data from car to attach results to session_valyes if its not soit will insert into session codes ';
-- //code_type will be (M "mechanical" , "T"  tires , "B" body and paint)
CREATE TABLE "codes" (
    "code_id" bigserial PRIMARY KEY,
    "car_brand_model_id" bigint NOT NULL,
    "code_name" character varying(200) NOT NULL,
    "vehicle_part" character varying(200) NOT NULL,
    "code_type" character varying(1) NOT NULL,
    "description" character varying(200) NOT NULL,
    "is_emergency" BOOLEAN
);
COMMENT ON COLUMN codes.code_type IS 'code_type will be (M "mechanical" , "T"  tires , "B" body and paint) ';

CREATE TABLE "session_values" (
    "session_value_id" bigserial PRIMARY KEY,
    "session_id" bigint NOT NULL,
    "value_key" character varying(50) NOT NULL,
    "value_data" character varying(50) NOT NULL
);
CREATE TABLE "session_codes" (
    "session_code_id" bigserial PRIMARY KEY,
    "session_id" bigint NOT NULL,
    "code_id"  bigint NOT NULL
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




ALTER TABLE
    "sessions"
ADD
    FOREIGN KEY ("car_id") REFERENCES "cars" ("car_id");

ALTER TABLE
    "session_values"
ADD
    FOREIGN KEY ("session_id") REFERENCES "sessions" ("session_id");



ALTER TABLE
    "session_codes"
ADD
    FOREIGN KEY ("session_id") REFERENCES "sessions" ("session_id");



ALTER TABLE
    "codes"
ADD
    FOREIGN KEY ("car_brand_model_id") REFERENCES "car_brand_models" ("car_brand_model_id");




ALTER TABLE
    "session_codes"
ADD
    FOREIGN KEY ("code_id") REFERENCES "codes" ("code_id");
