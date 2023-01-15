CREATE SEQUENCE IF NOT EXISTS account_id;
CREATE SEQUENCE IF NOT EXISTS pocket_id;
CREATE SEQUENCE IF NOT EXISTS txn_id;
CREATE SEQUENCE IF NOT EXISTS user_id;

CREATE TABLE IF NOT EXISTS "users" (
    "id" int4 NOT NULL DEFAULT nextval('user_id'::regclass),
    "username" VARCHAR(30) NOT NULL UNIQUE,
    "password" VARCHAR(30) NOT NULL,
    "create_dtm" TIMESTAMP DEFAULT now(),
    "update_dtm" TIMESTAMP DEFAULT now(),
    PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS index_username ON users (username);

CREATE TABLE IF NOT EXISTS "accounts" (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0, 
    "name" TEXT  NULL DEFAULT '',
    "user_id" int4 NOT NULL,
    "create_dtm" TIMESTAMP DEFAULT now(),
    "update_dtm" TIMESTAMP DEFAULT now(),
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "cloud_pocket" (
    "id" int4 NOT NULL DEFAULT nextval('pocket_id'::regclass),
    "name" TEXT NOT NULL,
    "account_id" int4 NOT NULL,
    "create_dtm" TIMESTAMP DEFAULT now(),
    "update_dtm" TIMESTAMP DEFAULT now(),
    "delete_dtm" TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "txn" (
    "id" int4 NOT NULL DEFAULT nextval('txn_id'::regclass),
    "timestamp" TIMESTAMP NOT NULL DEFAULT now(),
    "amount" NUMERIC NOT NULL,
    "note" VARCHAR NULL,
    "sender" int4 NOT NULL,
    "receiver" int4 NOT NULL,
    PRIMARY KEY ("id") 
);
