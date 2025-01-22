BEGIN;

-- Insert sample data into `users`
INSERT INTO users (id, preferred_username, given_name, family_name, email, created_at, updated_at)
VALUES
    ('548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'user1', 'User', 'One', 'user1@example.com', NOW(), NOW()),
    ('c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'user2', 'User', 'Two', 'user2@example.com', NOW(), NOW()),
    ('a4dfe207-f3bc-4710-8ba4-d9f69a45c093', 'user3', 'User', 'Three', 'user3@example.com', NOW(), NOW());

-- Insert sample data into `categories`
INSERT INTO categories (id, name, description, created_at, updated_at)
VALUES
    ('08d4b7cf-5618-42fa-943b-854e00e65d22', 'Fiction', 'Category for fictional books', NOW(), NOW()),
    ('7cc92a3d-8f39-414f-80c7-49b4c1201bd7', 'Science', 'Category for science books', NOW(), NOW());

-- Insert sample data into `authors`
INSERT INTO authors (id, first_name, last_name, nationality, birth_date, death_date, bio, website, created_at, updated_at)
VALUES
    ('e2435d88-dc2d-4c18-9068-85c3b5d885cd', 'Author', 'One', 'American', '1980-01-01', NULL, 'Bio of Author One', 'https://authorone.com', NOW(), NOW()),
    ('f9184633-798d-4b4f-9c4b-e01432683418', 'Author', 'Two', 'British', '1975-06-15', NULL, 'Bio of Author Two', NULL, NOW(), NOW()),
    ('3c8bc61f-def5-485b-a39a-0621e75f5914', 'Author', 'Three', 'Canadian', '1985-03-22', NULL, 'Bio of Author Three', NULL, NOW(), NOW()),
    ('36128715-c801-40ab-8a37-9b192af68fe2', 'Author', 'Four', 'French', '1990-07-19', NULL, 'Bio of Author Four', NULL, NOW(), NOW()),
    ('69c56de0-80f2-4cb5-a117-e3625ef01e93', 'Author', 'Five', 'German', '1982-11-03', NULL, 'Bio of Author Five', NULL, NOW(), NOW()),
    ('3435b61f-bbd1-47cb-9937-ec54d61e54f2', 'Author', 'Six', 'Italian', '1978-05-30', NULL, 'Bio of Author Six', NULL, NOW(), NOW()),
    ('e0583649-48bb-4f8c-804e-38d5fb7eec5e', 'Author', 'Seven', 'Spanish', '1992-09-17', NULL, 'Bio of Author Seven', NULL, NOW(), NOW()),
    ('71afdc7f-ae56-4224-9dd3-49fdf8ffa211', 'Author', 'Eight', 'Japanese', '1983-12-25', NULL, 'Bio of Author Eight', NULL, NOW(), NOW()),
    ('3f0651ec-cc6c-4817-ae60-8ab8ac8d6d62', 'Author', 'Nine', 'Indian', '1987-08-11', NULL, 'Bio of Author Nine', NULL, NOW(), NOW()),
    ('7c4e2e45-361a-4986-98bb-0bf5a6cd8906', 'Author', 'Ten', 'Russian', '1995-04-05', NULL, 'Bio of Author Ten', NULL, NOW(), NOW()),
    ('1dfa658d-3215-4cde-af12-e5595ae34167', 'Author', 'Eleven', 'Chinese', '1981-06-28', NULL, 'Bio of Author Eleven', NULL, NOW(), NOW()),
    ('2a9cbd25-c56a-4761-b708-beaa675f638d', 'Author', 'Twelve', 'Korean', '1976-10-20', NULL, 'Bio of Author Twelve', NULL, NOW(), NOW()),
    ('f20ba955-3304-4432-933c-62b5d17b4bc5', 'Author', 'Thirteen', 'Brazilian', '1991-03-14', NULL, 'Bio of Author Thirteen', NULL, NOW(), NOW()),
    ('812fbf9b-1e96-44d6-b6bf-0cbe50c05b57', 'Author', 'Fourteen', 'Australian', '1979-07-09', NULL, 'Bio of Author Fourteen', NULL, NOW(), NOW()),
    ('47082247-f6dd-4848-80b3-ec6bcb57dbea', 'Author', 'Fifteen', 'Mexican', '1986-02-02', NULL, 'Bio of Author Fifteen', NULL, NOW(), NOW()),
    ('9e5a9230-b394-4601-bb4e-8eb079a189bf', 'Author', 'Sixteen', 'Dutch', '1993-11-18', NULL, 'Bio of Author Sixteen', NULL, NOW(), NOW()),
    ('ed67da24-3da2-4c35-9ae8-64601ba2d843', 'Author', 'Seventeen', 'Swedish', '1984-09-13', NULL, 'Bio of Author Seventeen', NULL, NOW(), NOW()),
    ('02048d39-6cd6-4940-88f3-64d1bdf88d96', 'Author', 'Eighteen', 'Norwegian', '1974-01-26', NULL, 'Bio of Author Eighteen', NULL, NOW(), NOW()),
    ('9c3475dc-fb69-400b-9fa7-56ba9d3871b4', 'Author', 'Nineteen', 'Finnish', '1989-12-07', NULL, 'Bio of Author Nineteen', NULL, NOW(), NOW()),
    ('f1b5d5d8-c105-4670-a855-d61722881103', 'Author', 'Twenty', 'Portuguese', '1996-05-31', NULL, 'Bio of Author Twenty', NULL, NOW(), NOW()),
    ('6cc39622-8814-4e76-8a2f-b7393e77c9e7', 'Author', 'Twentyone', 'American', '1994-02-28', NULL, 'Bio of Author Twentyone', NULL, NOW(), NOW()),
    ('4a8f8fdd-97be-4d1c-a132-08191d3f9134', 'Author', 'Twentytwo', 'British', '1988-10-19', NULL, 'Bio of Author Twentytwo', NULL, NOW(), NOW()),
    ('1a4fbd47-3950-475b-bcd9-24f4a8b7c5a5', 'Author', 'Twentythree', 'Canadian', '1986-04-12', NULL, 'Bio of Author Twentythree', NULL, NOW(), NOW()),
    ('aa29de02-9262-429d-b7e6-73b12a703842', 'Author', 'Twentyfour', 'French', '1992-06-05', NULL, 'Bio of Author Twentyfour', NULL, NOW(), NOW()),
    ('cd4c2a47-b601-477a-9181-cd84ad47fd8e', 'Author', 'Twentyfive', 'German', '1984-09-29', NULL, 'Bio of Author Twentyfive', NULL, NOW(), NOW()),
    ('a3519de1-b7cc-49e5-9b68-e1a24fe49d96', 'Author', 'Twentysix', 'Italian', '1987-12-30', NULL, 'Bio of Author Twentysix', NULL, NOW(), NOW()),
    ('5a3f1bda-c17f-4025-8ef7-e93d7f3a907d', 'Author', 'Twentyseven', 'Spanish', '1989-03-23', NULL, 'Bio of Author Twentyseven', NULL, NOW(), NOW()),
    ('ed19e77b-cd4b-4ea5-b548-0731b591cbd9', 'Author', 'Twentyeight', 'Japanese', '1985-07-04', NULL, 'Bio of Author Twentyeight', NULL, NOW(), NOW()),
    ('125d2588-31b1-4f52-bcdc-379fa08e89c1', 'Author', 'Twentynine', 'Indian', '1993-09-19', NULL, 'Bio of Author Twentynine', NULL, NOW(), NOW()),
    ('019ab694-b1bc-45a2-853d-bd9a473ebd3c', 'Author', 'Thirty', 'Russian', '1980-11-04', NULL, 'Bio of Author Thirty', NULL, NOW(), NOW()),
    ('6de4c070-0a8d-44d7-8ffb-55a777e0f4f3', 'Author', 'Thirtyone', 'Chinese', '1990-03-17', NULL, 'Bio of Author Thirtyone', NULL, NOW(), NOW()),
    ('5e6e5930-d9b0-4967-b442-3df7e6984e1c', 'Author', 'Thirtytwo', 'Korean', '1988-02-22', NULL, 'Bio of Author Thirtytwo', NULL, NOW(), NOW()),
    ('11d62c64-7d39-47c9-b725-7a3445682056', 'Author', 'Thirtythree', 'Brazilian', '1986-05-30', NULL, 'Bio of Author Thirtythree', NULL, NOW(), NOW()),
    ('fbce9873-bf57-4a2b-bcd1-8853d957e478', 'Author', 'Thirtyfour', 'Australian', '1992-09-02', NULL, 'Bio of Author Thirtyfour', NULL, NOW(), NOW()),
    ('8485b5e4-3d1b-4022-bd29-11512ae3fa34', 'Author', 'Thirtyfive', 'Mexican', '1987-11-10', NULL, 'Bio of Author Thirtyfive', NULL, NOW(), NOW()),
    ('8725d975-1956-46c8-b4a5-bf3f914bd8f8', 'Author', 'Thirtysix', 'Dutch', '1994-10-14', NULL, 'Bio of Author Thirtysix', NULL, NOW(), NOW()),
    ('759a9a61-ef6d-47d1-a3a5-1a134d7fe22c', 'Author', 'Thirtyseven', 'Swedish', '1991-12-29', NULL, 'Bio of Author Thirtyseven', NULL, NOW(), NOW()),
    ('27a06f17-9351-42c0-8249-58c8da38fd2e', 'Author', 'Thirtyeight', 'Norwegian', '1995-04-20', NULL, 'Bio of Author Thirtyeight', NULL, NOW(), NOW()),
    ('4d9b0b52-b8b7-44c9-b118-1ea8e601cb5e', 'Author', 'Thirtynine', 'Finnish', '1989-08-15', NULL, 'Bio of Author Thirtynine', NULL, NOW(), NOW()),
    ('2bdb34b5-6ffb-46a5-b99f-bbe4cf529b8e', 'Author', 'Forty', 'Portuguese', '1993-07-23', NULL, 'Bio of Author Forty', NULL, NOW(), NOW()),
    ('5a513651-8d1d-4648-b80e-3a56e6d735c6', 'Author', 'Fortyone', 'American', '1985-02-17', NULL, 'Bio of Author Fortyone', NULL, NOW(), NOW()),
    ('4d984ccd-c7e1-4f17-bc92-e4f8d2fa1217', 'Author', 'Fortytwo', 'British', '1990-06-23', NULL, 'Bio of Author Fortytwo', NULL, NOW(), NOW()),
    ('9e837cfb-b99a-484d-aeb4-e5f7ab57a8ea', 'Author', 'Fortythree', 'Canadian', '1987-01-12', NULL, 'Bio of Author Fortythree', NULL, NOW(), NOW()),
    ('8d9e8bc2-d014-4a64-b4b0-568b7e795fbc', 'Author', 'Fortyfour', 'French', '1983-05-28', NULL, 'Bio of Author Fortyfour', NULL, NOW(), NOW()),
    ('94e0d276-cf16-4ba5-a137-e63d3cfb9db5', 'Author', 'Fortyfive', 'German', '1986-10-15', NULL, 'Bio of Author Fortyfive', NULL, NOW(), NOW()),
    ('a056f2e7-fb7d-462d-9513-45d80f91f981', 'Author', 'Fortysix', 'Italian', '1989-08-02', NULL, 'Bio of Author Fortysix', NULL, NOW(), NOW()),
    ('e3990f02-978b-4423-a437-0bb07ac1f9f7', 'Author', 'Fortyseven', 'Spanish', '1992-12-01', NULL, 'Bio of Author Fortyseven', NULL, NOW(), NOW()),
    ('13d43956-48a9-450d-8f6f-cc0f1f474d1f', 'Author', 'Fortyeight', 'Japanese', '1990-04-18', NULL, 'Bio of Author Fortyeight', NULL, NOW(), NOW()),
    ('0218e409-dc3c-4743-b040-8ea380fa6700', 'Author', 'Fortynine', 'Indian', '1994-11-30', NULL, 'Bio of Author Fortynine', NULL, NOW(), NOW()),
    ('a77365fd-8fe7-4715-a768-19d62a0bffbc', 'Author', 'Fifty', 'Russian', '1985-03-12', NULL, 'Bio of Author Fifty', NULL, NOW(), NOW()),
    ('9a56b03b-c5d2-453f-99ae-3c0b3f3db9a1', 'Author', 'Fiftyone', 'Chinese', '1992-08-21', NULL, 'Bio of Author Fiftyone', NULL, NOW(), NOW()),
    ('056d0f80-fc76-45bb-836d-48b59f1d7249', 'Author', 'Fiftytwo', 'Korean', '1980-02-17', NULL, 'Bio of Author Fiftytwo', NULL, NOW(), NOW()),
    ('b4d56a71-13b9-45b1-a6ad-fab4fa82efeb', 'Author', 'Fiftythree', 'Brazilian', '1986-09-08', NULL, 'Bio of Author Fiftythree', NULL, NOW(), NOW()),
    ('079b2295-961f-4230-b5c7-538b0a902a69', 'Author', 'Fiftyfour', 'Australian', '1993-07-05', NULL, 'Bio of Author Fiftyfour', NULL, NOW(), NOW()),
    ('22bb421e-1911-4a9e-8a89-b826e2f0c6a6', 'Author', 'Fiftyfive', 'Mexican', '1990-01-30', NULL, 'Bio of Author Fiftyfive', NULL, NOW(), NOW()),
    ('5f0bafc5-066b-4b92-89d1-6a9de75518fd', 'Author', 'Fiftysix', 'Dutch', '1987-05-18', NULL, 'Bio of Author Fiftysix', NULL, NOW(), NOW()),
    ('0be70532-f4fc-4a60-a2a0-450bca8b6879', 'Author', 'Fiftyseven', 'Swedish', '1984-10-22', NULL, 'Bio of Author Fiftyseven', NULL, NOW(), NOW()),
    ('5124f7a1-c8b9-450d-9b66-c9a2d4d5e9ac', 'Author', 'Fiftyeight', 'Norwegian', '1991-06-11', NULL, 'Bio of Author Fiftyeight', NULL, NOW(), NOW()),
    ('9b1a77d5-b3c5-4c70-baff-7ab71f8271be', 'Author', 'Fiftynine', 'Finnish', '1992-11-06', NULL, 'Bio of Author Fiftynine', NULL, NOW(), NOW()),
    ('04bfbef3-798a-4cfe-a02b-7cb81da2a98f', 'Author', 'Sixty', 'Portuguese', '1988-05-24', NULL, 'Bio of Author Sixty', NULL, NOW(), NOW()),
    ('ab839907-6e8b-4eb4-a149-e2b8db45087b', 'Author', 'Sixtyone', 'American', '1990-11-17', NULL, 'Bio of Author Sixtyone', NULL, NOW(), NOW()),
    ('c0e21c1b-b9b8-4ec5-9ed0-d7ccff99377d', 'Author', 'Sixtytwo', 'British', '1983-02-03', NULL, 'Bio of Author Sixtytwo', NULL, NOW(), NOW()),
    ('2ad51a7f-1a73-4fd7-9331-4694ac66efda', 'Author', 'Sixtythree', 'Canadian', '1985-09-12', NULL, 'Bio of Author Sixtythree', NULL, NOW(), NOW()),
    ('6cb64d7d-8d6e-41f1-a03f-8b0a289d0e02', 'Author', 'Sixtyfour', 'French', '1991-08-16', NULL, 'Bio of Author Sixtyfour', NULL, NOW(), NOW()),
    ('38d9b670-e01b-4a5d-8327-238314460f7b', 'Author', 'Sixtyfive', 'German', '1992-04-07', NULL, 'Bio of Author Sixtyfive', NULL, NOW(), NOW()),
    ('fa07f1db-1e6b-4a7c-bfb1-c201d12b5351', 'Author', 'Sixtysix', 'Italian', '1989-11-02', NULL, 'Bio of Author Sixtysix', NULL, NOW(), NOW()),
    ('d6364020-6250-4633-b7a5-4dbd9d74e79e', 'Author', 'Sixtyseven', 'Spanish', '1984-10-30', NULL, 'Bio of Author Sixtyseven', NULL, NOW(), NOW()),
    ('c3ffbd64-bc2e-46fa-a3a6-b8c605e7e66d', 'Author', 'Sixtyeight', 'Japanese', '1993-03-19', NULL, 'Bio of Author Sixtyeight', NULL, NOW(), NOW()),
    ('db7e6c89-f9ac-4299-8c07-6e1f5e5a67f1', 'Author', 'Sixtynine', 'Indian', '1982-07-14', NULL, 'Bio of Author Sixtynine', NULL, NOW(), NOW()),
    ('207cf65f-fd22-4f36-9086-c3a402f676d2', 'Author', 'Seventy', 'Russian', '1994-10-01', NULL, 'Bio of Author Seventy', NULL, NOW(), NOW()),
    ('0783e7d5-64b0-4d53-89c4-f54f9f9c5b3b', 'Author', 'Seventyone', 'Chinese', '1987-03-28', NULL, 'Bio of Author Seventyone', NULL, NOW(), NOW()),
    ('2db63a82-8d7f-4b9e-91b0-83a3487d84f7', 'Author', 'Seventytwo', 'Korean', '1986-08-12', NULL, 'Bio of Author Seventytwo', NULL, NOW(), NOW()),
    ('3820cf75-5d21-42f5-9729-d3b5479b91b1', 'Author', 'Seventythree', 'Brazilian', '1990-02-04', NULL, 'Bio of Author Seventythree', NULL, NOW(), NOW()),
    ('998c1327-d758-48d7-9b9e-01a3a4e056d0', 'Author', 'Seventyfour', 'Australian', '1985-01-15', NULL, 'Bio of Author Seventyfour', NULL, NOW(), NOW()),
    ('c04e7d1c-5d31-4659-9ff7-6c08b60c8706', 'Author', 'Seventyfive', 'Mexican', '1984-06-25', NULL, 'Bio of Author Seventyfive', NULL, NOW(), NOW()),
    ('6fa500e3-fd12-45a5-a8f6-221c8a5b4d8b', 'Author', 'Seventysix', 'Dutch', '1992-05-07', NULL, 'Bio of Author Seventysix', NULL, NOW(), NOW()),
    ('9117bb1d-08a1-4de4-950f-23d0b65b3df0', 'Author', 'Seventyseven', 'Swedish', '1983-11-12', NULL, 'Bio of Author Seventyseven', NULL, NOW(), NOW()),
    ('a809d7d3-bb0a-4b72-9379-617b2d2347fc', 'Author', 'Seventyeight', 'Norwegian', '1990-09-22', NULL, 'Bio of Author Seventyeight', NULL, NOW(), NOW()),
    ('d7b539d7-d2d4-47d7-9f98-d54f94893e1f', 'Author', 'Seventynine', 'Finnish', '1985-11-04', NULL, 'Bio of Author Seventynine', NULL, NOW(), NOW()),
    ('6899c079-5619-4067-9260-bbff3627e1b2', 'Author', 'Eighty', 'Portuguese', '1993-07-09', NULL, 'Bio of Author Eighty', NULL, NOW(), NOW());

-- Insert sample data into `books`
INSERT INTO books (id, title, year, author_id, category_id, total_copies, language, created_at, updated_at)
VALUES
    ('27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', 'Book One', 2020, 'e2435d88-dc2d-4c18-9068-85c3b5d885cd', '08d4b7cf-5618-42fa-943b-854e00e65d22', 5, 'English', NOW(), NOW()),
    ('bc582d8b-cbeb-45e6-8b0b-c89e871b5833', 'Book Two', 2018, 'f9184633-798d-4b4f-9c4b-e01432683418', '7cc92a3d-8f39-414f-80c7-49b4c1201bd7', 3, 'French', NOW(), NOW()),
    ('a1c61255-3547-4c6e-b3ba-f0ddf2763791', 'Book Three', 2022, 'e2435d88-dc2d-4c18-9068-85c3b5d885cd', NULL, 2, 'Spanish', NOW(), NOW());

-- Insert sample data into `loans`
INSERT INTO loans (id, user_id, book_id, start_date, due_date, return_date, status, created_at, updated_at)
VALUES
    ('dd69a43a-f51f-4f03-9302-c96f13b6b90b', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', '2024-01-01', '2024-01-15', NULL, 'active', NOW(), NOW()),
    ('b63c0269-df60-49cb-b815-685d8a587d6a', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', '2024-02-01', '2024-02-15', '2024-02-10', 'completed', NOW(), NOW());

-- Insert sample data into `reservations`
INSERT INTO reservations (id, user_id, book_id, created_at, expiry_date, updated_at)
VALUES
    ('ecdc1e64-8be3-4857-bf2f-c5a54d79f76a', 'a4dfe207-f3bc-4710-8ba4-d9f69a45c093', 'a1c61255-3547-4c6e-b3ba-f0ddf2763791', NOW(), '2025-03-01', NOW()),
    ('dd97be09-9467-4b22-91a8-f3c4bb743cc7', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', NOW(), '2025-02-20', NOW());

-- Insert sample data into `ratings`
INSERT INTO ratings (id, user_id, book_id, content, value, created_at, updated_at)
VALUES
    ('6bded01c-0207-42ed-8b09-5455b8167e25', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', 'Great book!', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d7139641d', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', 'Informative but hard to read.', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e21', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', '3', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396412', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', '4', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e23', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', '5', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396414', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', '6', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e26', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', '7', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396417', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', '8', 3, NOW(), NOW());

COMMIT;
