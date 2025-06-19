-- DOWN Migration
-- Dropping indexes
DROP INDEX IF EXISTS idx_products_product_name;
DROP INDEX IF EXISTS idx_products_deleted_at;
DROP INDEX IF EXISTS idx_categories_deleted_at;
DROP INDEX IF EXISTS idx_product_variants_deleted_at;
DROP INDEX IF EXISTS idx_products_active;
DROP INDEX IF EXISTS idx_categories_name_parent;
DROP INDEX IF EXISTS idx_product_variants_product_id;
DROP INDEX IF EXISTS idx_product_variants_attributes;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_user_id;
DROP INDEX IF EXISTS idx_product_variants_stock;

-- Dropping triggers
DROP TRIGGER IF EXISTS update_categories_timestamp ON categories;
DROP TRIGGER IF EXISTS update_products_timestamp ON products;
DROP TRIGGER IF EXISTS update_product_variants_timestamp ON product_variants;

-- Dropping function
DROP FUNCTION IF EXISTS update_timestamp;

-- Dropping tables
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS discounts;
DROP TABLE IF EXISTS price_history;
DROP TABLE IF EXISTS product_variants;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

-- Dropping extension
DROP EXTENSION IF EXISTS citext;