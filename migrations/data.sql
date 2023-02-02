INSERT INTO User (email, firstname, lastname, role, password)
VALUES
    ('owner@gmail.com', 'Owner', 'Propriétaire', 'owner', '123456'),
    ('admin@gmail.com', 'Admin', 'Administrateur', 'admin', '123456'),
    ('customer@gmail.com', 'Customer', 'Client', 'customer', '123456'),
    ('customer2@gmail.com', 'Customer2', 'Client2', 'customer', '123456'),
    ('owner2@gmail.com', 'Owner2', 'Propriétaire2', 'owner', '123456');

INSERT INTO Shop (name, description, owner_id)
VALUES
    ('Dentiste', 'Le meilleur dentiste de la capital', (SELECT id FROM User ORDER BY RAND() LIMIT 1)),
    ('Barber', 'Coiffeu afro', (SELECT id FROM User ORDER BY RAND() LIMIT 1));

INSERT INTO Service (id, name, description, price, duration, shop_id)
VALUES
    ('Blanchiment dentaire', 'Blanchiment dentaire de qualité', 150.00, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('Détartrage', 'Lorem ipsum', 30.99, 60,(SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('Coupe homme', 'Cheveux + barbe', 20.00, 60,(SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('Coupe barbe', 'Barbe', 10.00, 60, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1));

INSERT INTO Reservation (id, start, service_id, user_id)
VALUES
    ('70e8315c-45dd-480a-bf87-be2e99534afa', NOW(), 'a22ec672-a133-4427-ae3b-a085dbb7a5e3', (SELECT id FROM User LIMIT 1)),
    ('8f3d977a-b5ca-4532-bb81-36d1e56819e3', NOW(), '683569fb-dd2c-4c3e-bb52-349f767adda1', (SELECT id FROM User LIMIT 1));

INSERT INTO Hour (id, start, end, day, shop_id)
VALUES
    ('c20311b5-00b8-4dd3-b888-dd9b8f9cac9b', TIME('08:00:00'), TIME('18:00:00'), 6, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('45c866b8-8485-40c3-9b85-e9319b115298', TIME('08:00:00'), TIME('18:00:00'), 7, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('6cf1cdba-461d-465e-a38f-37a5458c796d', TIME('10:00:00'), TIME('20:00:00'), 6, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1)),
    ('c5684e7f-1705-4434-a5ad-d438be4773d8', TIME('10:00:00'), TIME('20:00:00'), 7, (SELECT id FROM Shop ORDER BY RAND() LIMIT 1));
