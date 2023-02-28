use planigo;

DELIMITER //
CREATE FUNCTION IF NOT EXISTS slugify(input_string VARCHAR(255))
    RETURNS VARCHAR(255)
BEGIN
    RETURN LOWER(REPLACE(input_string, ' ', '-'));
END;

CREATE FUNCTION IF NOT EXISTS new_uuid()
    RETURNS VARCHAR(36) CHARSET utf8
BEGIN
    RETURN UUID();
END;

CREATE TRIGGER IF NOT EXISTS `insert_user`
    BEFORE INSERT
    ON `User`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_store`
    BEFORE INSERT
    ON `Shop`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_service`
    BEFORE INSERT
    ON `Service`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_reservation`
    BEFORE INSERT
    ON `Reservation`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_hours`
    BEFORE INSERT
    ON `Hour`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_category`
    BEFORE INSERT
    ON `Category`
    FOR EACH ROW SET NEW.id = new_uuid();

CREATE TRIGGER IF NOT EXISTS `insert_category_slug`
    BEFORE INSERT
    ON `Category`
    FOR EACH ROW SET NEW.slug = slugify(NEW.name);

CREATE TRIGGER IF NOT EXISTS `insert_shop_slug`
    BEFORE INSERT
    ON `Shop`
    FOR EACH ROW SET NEW.slug = slugify(NEW.name);

//
DELIMITER ;
