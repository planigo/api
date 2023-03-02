use planigo;

-- Users
INSERT INTO User (email, firstname, lastname, role, password, is_email_verified)
VALUES ('john@example.com', 'John', 'Doe', 'admin', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('jane@example.com', 'Jane', 'Doe', 'owner', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('jim@example.com', 'Jim', 'Smith', 'customer', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1);

-- Categories
INSERT INTO Category (name)
VALUES ('Coiffure'),
       ('Soins des ongles'),
       ('Cosmétiques'),
       ('Spa'),
       ('Physiothérapie'),
       ('Chiropractie'),
       ('Acupuncture'),
       ('Massothérapie'),
       ('Dentisterie'),
       ('Soins dentaires'),
       ('Hygiène dentaire'),
       ('Esthéticienne');

-- Shops
INSERT INTO Shop (name, description, category_id, owner_id)
VALUES ('John\'s Hair Salon', 'Nous proposons un large éventail de services de coiffure.', (SELECT id FROM Category ORDER BY RAND() LIMIT 1), (SELECT id FROM User ORDER BY RAND() LIMIT 1)),
       ('Jane\'s Nail Salon', 'Nous proposons des services professionnels de soins des ongles pour les hommes et les femmes.', (SELECT id FROM Category ORDER BY RAND() LIMIT 1), (SELECT id FROM User ORDER BY RAND() LIMIT 1)),
       ('Jim\'s Massage Spa', 'Détendez et rajeunissez votre corps et votre esprit grâce à nos services de massothérapie.', (SELECT id FROM Category ORDER BY RAND() LIMIT 1),(SELECT id FROM User ORDER BY RAND() LIMIT 1)),
       ('Infinite Beauty', 'Détendez et rajeunissez votre corps et votre esprit grâce à nos services.', (SELECT id FROM Category ORDER BY RAND() LIMIT 1), (SELECT id FROM User ORDER BY RAND() LIMIT 1));

-- Hours
INSERT INTO Hour (start, end, day, shop_id)
VALUES (TIME('08:00:00'), TIME('18:00:00'), 1, (SELECT id FROM Shop ORDER BY name LIMIT 1)),
       (TIME('08:00:00'), TIME('18:00:00'), 2, (SELECT id FROM Shop ORDER BY name LIMIT 1)),
       (TIME('08:00:00'), TIME('18:00:00'), 3, (SELECT id FROM Shop ORDER BY name LIMIT 1)),
       (TIME('08:00:00'), TIME('18:00:00'), 4, (SELECT id FROM Shop ORDER BY name LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 5, (SELECT id FROM Shop ORDER BY name LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 1, (SELECT id FROM Shop ORDER BY name DESC LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 2, (SELECT id FROM Shop ORDER BY name DESC LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 3, (SELECT id FROM Shop ORDER BY name DESC LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 4, (SELECT id FROM Shop ORDER BY name DESC LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 5, (SELECT id FROM Shop ORDER BY name DESC LIMIT 1));

-- Services
INSERT INTO Service (name, description, price, duration, shop_id)
VALUES ('Coiffure', 'Coupe de cheveux', 50.00, 60, (SELECT id FROM Shop WHERE name LIKE 'John%' LIMIT 1)),
       ('Soins des ongles', 'Manucure et pédicure', 40.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jane%' LIMIT 1)),
       ('Cosmétiques', 'Soins de beauté pour le visage', 80.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Spa', 'Massages et soins de relaxation', 120.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Physiothérapie', 'Rééducation fonctionnelle', 80.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Chiropractie', 'Soins pour les problèmes de dos', 100.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Acupuncture', 'Médecine traditionnelle chinoise', 90.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Massothérapie', 'Massages thérapeutiques', 70.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jim%' LIMIT 1)),
       ('Dentisterie', 'Soins dentaires courants', 120.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jim%' LIMIT 1)),
       ('Hygiène dentaire', 'Soins préventifs et nettoyage', 80.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jim%' LIMIT 1)),
       ('Esthéticienne', 'Soins de beauté pour le corps', 60.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jim%' LIMIT 1)),
       ('Massothérapie', 'Massages de relaxation', 70.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Jim%' LIMIT 1)),
       ('Soins dentaires', 'Soins pour les problèmes dentaires', 150.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1)),
       ('Hygiène dentaire', 'Soins préventifs et nettoyage', 80.00, 60, (SELECT id FROM Shop WHERE name LIKE 'John%' LIMIT 1)),
       ('Coiffure', 'Coloration des cheveux', 80.00, 60, (SELECT id FROM Shop WHERE name LIKE 'John%' LIMIT 1)),
       ('Esthéticienne', 'Soins de beauté pour le visage', 60.00, 60, (SELECT id FROM Shop WHERE name LIKE 'Infinite%' LIMIT 1));

-- Reservations
INSERT INTO Reservation (start, service_id, user_id)
VALUES (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1));
