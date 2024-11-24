-- 2023110803_create_products_table.up.sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name_uz VARCHAR(255),
    name_ru VARCHAR(255),
    name_en VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    description_uz TEXT,
    description_ru TEXT,
    description_en TEXT,
    category_id INT NOT NULL,
    brand_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE CASCADE
);
