create table if not exists `category` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    icon VARCHAR(255) NULL,
    category_name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    notes TEXT NULL,
    is_aktif ENUM('0','1') NOT NULL DEFAULT '1',
    parent_id INTEGER NULL,
    created_at TIMESTAMP DEFAULT (now()),
    updated_at TIMESTAMP DEFAULT (now())
);


-- Create foreign key constraints with table "product";
ALTER TABLE `product` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

