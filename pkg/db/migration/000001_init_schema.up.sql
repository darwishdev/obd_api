
CREATE TABLE "branches" (
    "branch_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "location" character varying(200) NOT NULL,
    "created_at" timestamptz NOT NULL  DEFAULT NOW(),
    "deleted_at" timestamptz
);
CREATE TABLE "users" (
    "user_id" bigserial PRIMARY KEY,
    "name" character varying(200) NOT NULL,
    "phone" character varying(200) UNIQUE NOT NULL,
    "email" character varying(200) UNIQUE NOT NULL,
    "password" character varying(200) NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00'),  
    "created_at" timestamptz NOT NULL  DEFAULT NOW(),
    "deleted_at" timestamptz

);


CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
    "product_id" bigserial PRIMARY KEY,
    "name" character varying(200) CHECK (length(name) >= 2) NOT NULL,
    "code" character varying(20) CHECK (length(code) >= 2) UNIQUE NOT NULL,
    "price" real CHECK (price > 0) NOT NULL,
    "description" text,
    "created_at" timestamptz NOT NULL  DEFAULT NOW(),
    "deleted_at" timestamptz
);



CREATE TABLE "orders" (
    "order_id" bigserial PRIMARY KEY,
    "order_number" character varying(20) NOT NULL,
    "total_amount" real NOT NULL,
    "user_id"  bigint NOT NULL,
    "branch_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL  DEFAULT NOW()
);


CREATE TABLE "order_products" (
    "order_products_id" bigserial PRIMARY KEY,
    "order_id"  bigint NOT NULL,
    "product_id"  bigint NOT NULL,
    "qty" real NOT NULL,
    "unit_price" real NOT NULL ,
    "total_amount" real NOT NULL,
    "deleted_at" timestamptz
);


ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "orders" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("branch_id");
ALTER TABLE "order_products" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");
ALTER TABLE "order_products" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");