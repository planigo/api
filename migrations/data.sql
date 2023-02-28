use planigo;

INSERT INTO User (email, firstname, lastname, role, password, is_email_verified)
VALUES ('owner@gmail.com', 'Owner', 'Propriétaire', 'owner', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('admin@gmail.com', 'Admin', 'Administrateur', 'admin', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('customer@gmail.com', 'Customer', 'Client', 'customer', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('customer2@gmail.com', 'Customer2', 'Client2', 'customer', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1),
       ('owner2@gmail.com', 'Owner2', 'Propriétaire2', 'owner', '$2a$14$Rs4.s9vRidmhhcExUV0/Iu6ihdSXkw4yTClU.0vpmR.savUfjdq8W', 1);

INSERT INTO Category (slug, name)
VALUES ('', 'Coiffeur'),
       ('', 'Barber'),
       ('', 'Dentiste');

INSERT INTO Shop (name, description, category_id, owner_id)
VALUES ('Dentiste', 'Le meilleur dentiste de la capital', (SELECT id FROM Category ORDER BY RAND() LIMIT 1), (SELECT id FROM User ORDER BY RAND() LIMIT 1)),
       ('Barber', 'Coiffeu afro', (SELECT id FROM Category ORDER BY RAND() LIMIT 1), (SELECT id FROM User ORDER BY RAND() LIMIT 1));

INSERT INTO Service (name, description, price, duration, shop_id)
VALUES ('Blanchiment dentaire', 'Blanchiment dentaire de qualité', 150.00, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       ('Détartrage', 'Lorem ipsum', 30.99, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       ('Coupe homme', 'Cheveux + barbe', 20.00, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       ('Coupe barbe', 'Barbe', 10.00, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1));

INSERT INTO Reservation (start, service_id, user_id)
VALUES (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1)),
       (NOW(), (SELECT id FROM Service ORDER BY RAND() LIMIT 1), (SELECT id FROM User LIMIT 1));

INSERT INTO Hour (start, end, day, shop_id)
VALUES (TIME('08:00:00'), TIME('18:00:00'), 6, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       (TIME('08:00:00'), TIME('18:00:00'), 7, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 6, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
       (TIME('10:00:00'), TIME('20:00:00'), 7, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1));
