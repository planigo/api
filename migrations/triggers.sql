use planigo;

CREATE FUNCTION `slugify`(str VARCHAR(255)) RETURNS VARCHAR(255) DETERMINISTIC RETURN LOWER(REPLACE(str, ' ', '-'));

CREATE TRIGGER `insert_user`BEFORE INSERT ON `User` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_store`BEFORE INSERT ON `Shop` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_service`BEFORE INSERT ON `Service` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_reservation`BEFORE INSERT ON `Reservation` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_hours`BEFORE INSERT ON `Hour` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_category` BEFORE INSERT ON `Category` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_category_slug` BEFORE INSERT ON `Category` FOR EACH ROW SET NEW.slug = slugify(NEW.name);

CREATE TRIGGER `insert_shop_slug` BEFORE INSERT ON `Shop` FOR EACH ROW SET NEW.slug = slugify(NEW.name);