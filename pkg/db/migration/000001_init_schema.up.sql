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
