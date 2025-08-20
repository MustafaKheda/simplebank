ALTER TABLE IF EXISTS "accounts" DROP constraint "owner_currency_key";

ALTER TABLE IF EXISTS "accounts" DROP constraint "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";