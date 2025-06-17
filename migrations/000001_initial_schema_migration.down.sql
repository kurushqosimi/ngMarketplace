-- Drop comments
COMMENT ON TABLE categories IS NULL;
COMMENT ON TABLE products IS NULL;
COMMENT ON TABLE product_variants IS NULL;

-- Drop indexes
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

-- Drop triggers
DROP TRIGGER IF EXISTS update_categories_timestamp ON categories;
DROP TRIGGER IF EXISTS update_products_timestamp ON products;
DROP TRIGGER IF EXISTS update_product_variants_timestamp ON product_variants;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_timestamp;

-- Drop foreign key constraints
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS fk_categories_parent_id;
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS fk_products_category_id;
ALTER TABLE "product_variants" DROP CONSTRAINT IF EXISTS fk_product_variants_product_id;

-- Drop tables
DROP TABLE IF EXISTS product_variants;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

-- Drop citext extension (only if no other tables use it)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_attribute
        JOIN pg_class ON pg_attribute.attrelid = pg_class.oid
        JOIN pg_type ON pg_attribute.atttypid = pg_type.oid
        WHERE pg_type.typname = 'citext'
        AND pg_class.relname NOT IN ('products')
    ) THEN
        DROP EXTENSION IF EXISTS citext;
END IF;
END $$;