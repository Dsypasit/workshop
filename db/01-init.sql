CREATE SEQUENCE IF NOT EXISTS account_id;
CREATE SEQUENCE IF NOT EXISTS pocket_id;
CREATE SEQUENCE IF NOT EXISTS txn_id;

CREATE TABLE IF NOT EXISTS "accounts" (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "cloud_pocket" (
    "id" int4 NOT NULL DEFAULT nextval('pocket_id'::regclass),
    "name" TEXT NOT NULL,
    PRIMARY KEY ("id")
)

CREATE TABLE IF NOT EXISTS "txn" (
    "id" int4 NOT NULL DEFAULT nextval('txn_id'::regclass),
    "timestamp" TIMESTAMP NOT NULL,
    "amount" NUMERIC NOT NULL,
    "note" VARCHAR NOT NULL,
    "sender" int4 NOT NULL,
    "receiver" int4 NOT NULL,
    PRIMARY KEY ("id") 
)