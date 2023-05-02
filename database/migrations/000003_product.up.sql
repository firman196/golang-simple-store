create table if not exists `product` (
    sku VARCHAR(50) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    sallable_stock INTEGER NOT NULL DEFAULT 0,
    weight INTEGER NOT NULL DEFAULT 0,
    description TEXT NOT NULL,
    rating FLOAT(10,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT (now()),
    updated_at TIMESTAMP DEFAULT (now())
);

-- Create foreign key constraints with table "product";
ALTER TABLE `product_images` ADD FOREIGN KEY (`product_sku`) REFERENCES `product` (`sku`) ON DELETE CASCADE ON UPDATE CASCADE;
