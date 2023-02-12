CREATE DATABASE IF NOT EXISTS `planigo`;

CREATE TABLE IF NOT EXISTS `User`
(
    `id`                VARCHAR(36) PRIMARY KEY NOT NULL,
    `email`             VARCHAR(255) UNIQUE     NOT NULL,
    `firstname`         VARCHAR(255)            NOT NULL,
    `lastname`          VARCHAR(255)            NOT NULL,
    `role`              ENUM ('admin', 'owner', 'customer') DEFAULT 'customer',
    `password`          TEXT                    NOT NULL,
    `is_email_verified` BOOLEAN                             DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS `Category`
(
    `id`   VARCHAR(36) PRIMARY KEY NOT NULL,
    `slug` VARCHAR(255) UNIQUE     NOT NULL,
    `name` VARCHAR(255)            NOT NULL
);

CREATE TABLE IF NOT EXISTS `Shop`
(
    `id`          VARCHAR(36) PRIMARY KEY NOT NULL,
    `slug`        VARCHAR(255)            NOT NULL,
    `name`        VARCHAR(255)            NOT NULL,
    `description` TEXT,
    `owner_id`    VARCHAR(36),
    `category_id` VARCHAR(36),
    CONSTRAINT `fk_shop_category`
        FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`),
    CONSTRAINT `fk_shop_owner`
        FOREIGN KEY (`owner_id`) REFERENCES `User` (`id`)
);

CREATE TABLE IF NOT EXISTS `Service`
(
    `id`          VARCHAR(36) PRIMARY KEY NOT NULL,
    `name`        VARCHAR(255)            NOT NULL,
    `description` TEXT,
    `price`       FLOAT,
    `duration`    INTEGER DEFAULT 60,
    `shop_id`     VARCHAR(36),
    CONSTRAINT `fk_service_shop`
        FOREIGN KEY (`shop_id`) REFERENCES `Shop` (`id`)
);

CREATE TABLE IF NOT EXISTS `Reservation`
(
    `id`           VARCHAR(36) PRIMARY KEY NOT NULL,
    `start`        datetime,
    `service_id`   VARCHAR(36),
    `user_id`      VARCHAR(36),
    `is_cancelled` BOOLEAN DEFAULT FALSE,
    CONSTRAINT `fk_reservation_service`
        FOREIGN KEY (`service_id`) REFERENCES `Service` (`id`),
    CONSTRAINT `fk_reservation_user`
        FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
);

CREATE TABLE IF NOT EXISTS `Hour`
(
    `id`      VARCHAR(36) PRIMARY KEY NOT NULL,
    `start`   time,
    `end`     time,
    `day`     integer,
    `shop_id` VARCHAR(36),
    CONSTRAINT `fk_hours_shop`
        FOREIGN KEY (`shop_id`) REFERENCES `Shop` (`id`)
);
