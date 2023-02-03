use planigo;

CREATE TRIGGER `insert_user`BEFORE INSERT ON `User` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_store`BEFORE INSERT ON `Shop` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_service`BEFORE INSERT ON `Service` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_reservation`BEFORE INSERT ON `Reservation` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_hours`BEFORE INSERT ON `Hour` FOR EACH ROW SET NEW.id = UUID();

CREATE TRIGGER `insert_category` BEFORE INSERT ON `Category` FOR EACH ROW SET NEW.id = UUID();