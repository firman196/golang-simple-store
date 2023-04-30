create table if not exists `users` (
    id INTEGER PRIMARY key AUTO_INCREMENT,
    firstname VARCHAR(100) NOTE NULL,
    lastname VARCHAR(100) NULL,
    password TEXT NOT NULL,
    salt VARCHAR(100) NULL,
    address TEXT NULL,
    image VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    email_verified_at TIMESTAMP NULL,
    verivication_code VARCHAR(10) NULL, 
    phone_number VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at TIMESTAMP NOT NULL DEFAULT (now()),
    is_banned ENUM(0,1) NOT NULL DEFAULT 0,
    is_delete ENUM(0,1) NOT NULL DEFAULT 0,
    deleted_at TIMESTAMP NULL,
    deleted_by INTEGER NULL
);

