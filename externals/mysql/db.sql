CREATE DATABASE IF NOT EXISTS kudaki_rental DEFAULT COLLATE = utf8_general_ci;
CREATE USER IF NOT EXISTS 'kudaki_user' @'localhost' IDENTIFIED BY 'kudakirocks';
GRANT ALL PRIVILEGES ON kudaki_rental.* TO 'kudaki_user' @'localhost' WITH GRANT OPTION;
USE kudaki_rental;
CREATE TABLE IF NOT EXISTS carts(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(255) UNIQUE,
  `user_uuid` VARCHAR(255),
  `total_price` INT(20),
  `open` TINYINT(1)
);
CREATE TABLE IF NOT EXISTS cart_items(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(255) UNIQUE,
  `cart_uuid` VARCHAR(255) UNIQUE,
  `item_uuid` VARCHAR(255),
  `total_item` SMALLINT(5),
  `total_price` INT(20),
  `unit_price` INT(20),
  FOREIGN KEY(cart_uuid) REFERENCES carts(uuid) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS checkouts(
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uuid` VARCHAR(255) UNIQUE,
  `cart_uuid` VARCHAR(255) UNIQUE,
  `issued_at` DATETIME,
  FOREIGN KEY(cart_uuid) REFERENCES carts(uuid) ON DELETE CASCADE
);
