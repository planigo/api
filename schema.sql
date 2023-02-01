CREATE TABLE `User`
(
    `id`        varchar(36) PRIMARY KEY NOT NULL,
    `email`     varchar(255)            NOT NULL,
    `firstname` varchar(255)            NOT NULL,
    `lastname`  varchar(255)            NOT NULL,
    `role`      ENUM ('admin', 'seller', 'customer'),
    `password`  text                    NOT NULL
);

CREATE TABLE `Store`
(
    `id`          varchar(36) PRIMARY KEY NOT NULL,
    `name`        varchar(255)            NOT NULL,
    `description` text,
    `owner_id`    varchar(36),
    constraint `fk_store_owner`
        foreign key (`owner_id`) references `User` (`id`)
);

CREATE TABLE `Service`
(
    `id`          varchar(36) PRIMARY KEY NOT NULL,
    `name`        varchar(255)            NOT NULL,
    `description` text,
    `price`       float,
    `store_id`    varchar(36),
    `duration`    integer default 60,
    constraint `fk_service_store`
        foreign key (`store_id`) references `Store` (`id`)
);

CREATE TABLE `Reservation`
(
    `id`         varchar(36) PRIMARY KEY NOT NULL,
    `start`      datetime,
    `service_id` varchar(36),
    `user_id`    varchar(36),
    constraint `fk_reservation_service`
        foreign key (`service_id`) references `Service` (`id`),
    constraint `fk_reservation_user`
        foreign key (`user_id`) references `User` (`id`)
);

CREATE TABLE `Hours`
(
    `id`       varchar(36) PRIMARY KEY NOT NULL,
    `start`    time,
    `end`      time,
    `day`      integer,
    `store_id` varchar(36),
    constraint `fk_hours_store`
        foreign key (`store_id`) references `Store` (`id`)
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
    ON `Hours`
    FOR EACH ROW
BEGIN
    SET NEW.id = UUID();
END;