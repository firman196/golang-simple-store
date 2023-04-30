create table if not exists `product_images` ( 
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    product_id INTEGER NOT NULL,
    image TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create foreign key constraints with table "product";
ALTER TABLE `product_images` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

