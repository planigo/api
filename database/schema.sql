CREATE TABLE `User`
(
    `id`        VARCHAR(36) PRIMARY KEY NOT NULL,
    `email`     VARCHAR(255)            NOT NULL,
    `firstname` VARCHAR(255)            NOT NULL,
    `lastname`  VARCHAR(255)            NOT NULL,
    `role`      ENUM ('admin', 'owner', 'customer') DEFAULT 'customer',
    `password`  TEXT                    NOT NULL,
    `is_email_verified` BOOLEAN DEFAULT FALSE
);

CREATE TABLE `Store`
(
    `id`          VARCHAR(36) PRIMARY KEY NOT NULL,
    `name`        VARCHAR(255)            NOT NULL,
    `description` TEXT,
    `owner_id`    VARCHAR(36),
    CONSTRAINT `fk_store_owner`
        FOREIGN KEY (`owner_id`) REFERENCES `User` (`id`)
);

CREATE TABLE `Service`
(
    `id`          VARCHAR(36) PRIMARY KEY NOT NULL,
    `name`        VARCHAR(255)            NOT NULL,
    `description` TEXT,
    `price`       FLOAT,
    `store_id`    VARCHAR(36),
    `duration`    INTEGER DEFAULT 60,
    CONSTRAINT `fk_service_store`
        FOREIGN KEY (`store_id`) REFERENCES `Store` (`id`)
);

CREATE TABLE `Reservation`
(
    `id`         VARCHAR(36) PRIMARY KEY NOT NULL,
    `start`      datetime,
    `service_id` VARCHAR(36),
    `user_id`    VARCHAR(36),
    CONSTRAINT `fk_reservation_service`
        FOREIGN KEY (`service_id`) REFERENCES `Service` (`id`),
    CONSTRAINT `fk_reservation_user`
        FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
);

CREATE TABLE `Hour`
(
    `id`       VARCHAR(36) PRIMARY KEY NOT NULL,
    `start`    time,
    `end`      time,
    `day`      integer,
    `store_id` VARCHAR(36),
    CONSTRAINT `fk_hours_store`
        FOREIGN KEY (`store_id`) REFERENCES `Store` (`id`)
);

CREATE TRIGGER `insert_user`
    BEFORE INSERT
    ON `User`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;

CREATE TRIGGER `insert_store`
    BEFORE INSERT
    ON `Store`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;

CREATE TRIGGER `insert_service`
    BEFORE INSERT
    ON `Service`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;

CREATE TRIGGER `insert_reservation`
    BEFORE INSERT
    ON `Reservation`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;

CREATE TRIGGER `insert_hours`
    BEFORE INSERT
    ON `Hour`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;
