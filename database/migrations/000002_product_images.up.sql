create table if not exists `product_images` ( 
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    product_sku VARCHAR(50) NOT NULL,
    image TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at TIMESTAMP NOT NULL DEFAULT (now())
);



