-- First we need to drop the unique constraint "owner_currency_key"
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Drop the foreign key constraint fo the owner field
-- you can find the name constraint in the table plus info
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";
