INSERT INTO User (id, email, firstname, lastname, role, password)
VALUES
    ('8d74da88-10f0-4602-9f3d-65b88b231d68', 'admin@gmail.com', 'Admin', 'Administrateur', 'admin', '123456'),
    ('6f3e6db7-6185-4116-a89c-a06713e5fb62', 'customer@gmail.com', 'Customer', 'Client', 'customer', '123456'),
    ('b21872c6-0ba4-4a9d-ae86-4be7430f81a8', 'owner@gmail.com', 'Owner', 'Propriétaire', 'owner', '123456'),
    ('793aeb6b-78f5-44b0-bde9-f5e999010a4a', 'customer2@gmail.com', 'Customer2', 'Client2', 'customer', '123456'),
    ('37d7a2f2-10d2-44f7-b71b-d063109b2da5', 'owner2@gmail.com', 'Owner2', 'Propriétaire2', 'owner', '123456');

INSERT INTO Shop (id, name, description, owner_id)
VALUES
    ('83452549-f594-4317-a3b0-c72a626b73f7', 'Dentiste', 'Le meilleur dentiste de la capital', 'b21872c6-0ba4-4a9d-ae86-4be7430f81a8'),
    ('a9d9fa06-e0c1-481a-9b53-b0a58f6a74b2', 'Barber', 'Coiffeu afro', '37d7a2f2-10d2-44f7-b71b-d063109b2da5');

INSERT INTO Service (id, name, description, price, duration, shop_id)
VALUES
    ('5f96f1c0-0043-47c1-b360-650800baded7', 'Blanchiment dentaire', 'Blanchiment dentaire de qualité', 150.00, 60, '83452549-f594-4317-a3b0-c72a626b73f7'),
    ('683569fb-dd2c-4c3e-bb52-349f767adda1', 'Détartrage', 'Lorem ipsum', 30.99, 60, '83452549-f594-4317-a3b0-c72a626b73f7'),
    ('a22ec672-a133-4427-ae3b-a085dbb7a5e3', 'Coupe homme', 'Cheveux + barbe', 20.00, 60, 'a9d9fa06-e0c1-481a-9b53-b0a58f6a74b2'),
    ('97686418-6019-4d5c-bcd3-9dab32786340', 'Coupe barbe', 'Barbe', 10.00, 60, 'a9d9fa06-e0c1-481a-9b53-b0a58f6a74b2');

INSERT INTO Reservation (id, start, service_id, user_id)
VALUES
    ('70e8315c-45dd-480a-bf87-be2e99534afa', NOW(), 'a22ec672-a133-4427-ae3b-a085dbb7a5e3', '6f3e6db7-6185-4116-a89c-a06713e5fb62'),
    ('8f3d977a-b5ca-4532-bb81-36d1e56819e3', NOW(), '683569fb-dd2c-4c3e-bb52-349f767adda1', '793aeb6b-78f5-44b0-bde9-f5e999010a4a');

INSERT INTO Hour (id, start, end, day, shop_id)
VALUES
    ('c20311b5-00b8-4dd3-b888-dd9b8f9cac9b', TIME('08:00:00'), TIME('18:00:00'), 6, '83452549-f594-4317-a3b0-c72a626b73f7'),
    ('45c866b8-8485-40c3-9b85-e9319b115298', TIME('08:00:00'), TIME('18:00:00'), 7, '83452549-f594-4317-a3b0-c72a626b73f7'),
    ('6cf1cdba-461d-465e-a38f-37a5458c796d', TIME('10:00:00'), TIME('20:00:00'), 6, 'a9d9fa06-e0c1-481a-9b53-b0a58f6a74b2'),
    ('c5684e7f-1705-4434-a5ad-d438be4773d8', TIME('10:00:00'), TIME('20:00:00'), 7, 'a9d9fa06-e0c1-481a-9b53-b0a58f6a74b2');
