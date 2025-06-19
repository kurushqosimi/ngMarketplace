-- Удаление индексов
DROP INDEX IF EXISTS idx_product_variants_stock;
DROP INDEX IF EXISTS idx_products_user_id;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_product_variants_attributes;
DROP INDEX IF EXISTS idx_product_variants_product_id;
DROP INDEX IF EXISTS idx_categories_name_parent;
DROP INDEX IF EXISTS idx_products_active;
DROP INDEX IF EXISTS idx_product_variants_deleted_at;
DROP INDEX IF EXISTS idx_categories_deleted_at;
DROP INDEX IF EXISTS idx_products_deleted_at;
DROP INDEX IF EXISTS idx_products_product_name;

-- Удаление триггеров
DROP TRIGGER IF EXISTS update_product_variants_timestamp ON product_variants;
DROP TRIGGER IF EXISTS update_products_timestamp ON products;
DROP TRIGGER IF EXISTS update_categories_timestamp ON categories;

-- Удаление функции
DROP FUNCTION IF EXISTS update_timestamp;

-- Удаление внешних ключей
ALTER TABLE product_images DROP CONSTRAINT IF EXISTS product_images_product_id_fkey;
ALTER TABLE product_images DROP CONSTRAINT IF EXISTS product_images_variant_id_fkey;
ALTER TABLE discounts DROP CONSTRAINT IF EXISTS discounts_product_id_fkey;
ALTER TABLE discounts DROP CONSTRAINT IF EXISTS discounts_variant_id_fkey;
ALTER TABLE discounts DROP CONSTRAINT IF EXISTS discounts_category_id_fkey;
ALTER TABLE price_history DROP CONSTRAINT IF EXISTS price_history_product_id_fkey;
ALTER TABLE price_history DROP CONSTRAINT IF EXISTS price_history_variant_id_fkey;
ALTER TABLE product_variants DROP CONSTRAINT IF EXISTS product_variants_product_id_fkey;
ALTER TABLE products DROP CONSTRAINT IF EXISTS products_category_id_fkey;
ALTER TABLE categories DROP CONSTRAINT IF EXISTS categories_parent_id_fkey;

-- Удаление таблиц
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS discounts;
DROP TABLE IF EXISTS price_history;
DROP TABLE IF EXISTS product_variants;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

-- Удаление расширения citext
DROP EXTENSION IF EXISTS citext;