-- Drop triggers
DROP TRIGGER IF EXISTS update_products_translation_timestamp ON product_translations;
DROP TRIGGER IF EXISTS update_products_timestamp ON products;
DROP TRIGGER IF EXISTS update_categories_timestamp ON categories;

-- Drop functions
DROP FUNCTION IF EXISTS update_timestamp();

-- Drop indexes of product_translations
DROP INDEX IF EXISTS idx_product_translations_attributes;
DROP INDEX IF EXISTS idx_products_language;
DROP INDEX IF EXISTS idx_product_translations_search;
DROP INDEX IF EXISTS idx_product_id_language;

-- Drop table product_translations
DROP TABLE IF EXISTS product_translations;

-- Drop indexes of products
DROP INDEX IF EXISTS idx_products_user;
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_products_active;

-- Drop table products
DROP TABLE IF EXISTS products;

-- Drop indexes of categories
DROP INDEX IF EXISTS idx_categories_language;
DROP INDEX IF EXISTS idx_categories_name_search;
DROP INDEX IF EXISTS idx_categories_languages;
DROP INDEX IF EXISTS idx_categories_name_parent_lang;
DROP INDEX IF EXISTS idx_categories_deleted_at;

-- Drop table categories
DROP TABLE IF EXISTS categories;

-- Drop extension citext
DROP EXTENSION IF EXISTS citext;