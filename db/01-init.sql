CREATE SEQUENCE IF NOT EXISTS account_id;

CREATE TABLE "accounts" (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE TABLE "pockets" (
    "id" int4 NOT NULL DEFAULT nextval('pocket_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    "account_id" int4 NOT NULL
    "name" TEXT NOT NULL
    "currency" TEXT NOT NULL
    "category" TEXT NOT NULL
    PRIMARY KEY ("id")
)