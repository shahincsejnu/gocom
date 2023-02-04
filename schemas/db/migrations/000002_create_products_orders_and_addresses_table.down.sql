ALTER TABLE "users" DROP COLUMN "is_admin";

ALTER TABLE "orders" DROP CONSTRAINT "orders_product_id_fkey";

ALTER TABLE "orders" DROP CONSTRAINT "orders_user_id_fkey";

ALTER TABLE "orders" DROP CONSTRAINT "orders_address_id_fkey";

ALTER TABLE "addresses" DROP CONSTRAINT "addresses_user_id_fkey";

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS addresses;
