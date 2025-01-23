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
    ('7cc92a3d-8f39-414f-80c7-49b4c1201bd7', 'Science', 'Category for science books', NOW(), NOW()),
    ('f91d2b93-1a8a-4b8a-9e74-2d7d892d784c', 'History', 'Category for historical books', NOW(), NOW()),
    ('d72e6cbb-d748-4536-9952-7e62e243c213', 'Technology', 'Books related to technology and innovations', NOW(), NOW()),
    ('ecc9c34c-9b4d-4c4a-bc3d-68cb570ea23c', 'Self-Help', 'Books for personal development and self-improvement', NOW(), NOW()),
    ('29f5f9d3-2c36-4e62-88f3-06fc8a1ebee6', 'Business', 'Books on business strategies and entrepreneurship', NOW(), NOW()),
    ('3a9d26e6-d5aa-4413-903e-78a1c9ddf0db', 'Cooking', 'Recipes and culinary arts books', NOW(), NOW()),
    ('0bd2e615-72d8-4b55-82b2-53f47bcd4c31', 'Travel', 'Books about travel destinations and guides', NOW(), NOW()),
    ('12ebef67-99d2-4c15-a5b3-b3843f3e5b14', 'Philosophy', 'Books about philosophical thoughts and ideas', NOW(), NOW()),
    ('2d93b5c7-82e5-48a2-b5fd-1a1f7e2ed74b', 'Poetry', 'Collection of poems and lyrical literature', NOW(), NOW()),
    ('c8e3f2a1-6209-4d37-9a1d-df994eb5a8f3', 'Romance', 'Books about love and relationships', NOW(), NOW()),
    ('f0cb5e0a-0c95-43c4-81d3-6f97936c477d', 'Mystery', 'Detective and mystery novels', NOW(), NOW()),
    ('d4f212aa-8b3d-46a3-9d3e-3084c2831dc1', 'Horror', 'Books that explore horror and supernatural themes', NOW(), NOW()),
    ('8e6dbd90-3e52-498b-8230-91c8e758e327', 'Fantasy', 'Books set in fantastical worlds with magic and creatures', NOW(), NOW()),
    ('7f2126d3-01d1-4828-9e5a-78f45d857db5', 'Comics', 'Comic books and graphic novels', NOW(), NOW()),
    ('e5b2461b-62ff-43e0-90cb-4e89d8e57db4', 'Art', 'Books about art history and techniques', NOW(), NOW()),
    ('12d93b5c-22e5-47a2-b3fd-1a4f7e2ed76b', 'Photography', 'Books showcasing photography techniques and collections', NOW(), NOW()),
    ('fe13dabc-4239-42b2-912c-90c6218a5b54', 'Sports', 'Books about various sports and fitness topics', NOW(), NOW()),
    ('0f5c5e24-3e5a-4b55-b2c2-55f45fcd4d12', 'Science Fiction', 'Books exploring futuristic science and technology', NOW(), NOW()),
    ('9d5e2b15-1e94-4a1b-bc2e-1f3926d31d23', 'Biography', 'Biographies and autobiographies of famous people', NOW(), NOW()),
    ('2bd91e5c-72f8-4d55-92b2-67f45bcd4c99', 'Health', 'Books on health, wellness, and medical topics', NOW(), NOW()),
    ('cd4f5f12-8b3d-42a3-9d3e-3084c2837dd5', 'Children', 'Books for kids and early education', NOW(), NOW()),
    ('82c3f1b2-6209-4d37-9a3d-df994eb5a1b3', 'Psychology', 'Books about human mind and behavior', NOW(), NOW()),
    ('ef5a6b1a-5e6d-43b2-9e92-2c7d832d745d', 'Politics', 'Books related to political studies and government', NOW(), NOW()),
    ('b1e5c3a1-6229-42b2-9e5d-df794eb5a2f3', 'Religion', 'Books about different religions and spiritual paths', NOW(), NOW()),
    ('7a5c2e1b-92f8-4c55-82b2-59f45ecd4c88', 'Music', 'Books about music theory and musicians', NOW(), NOW()),
    ('5f2a3c1e-6b2d-48e2-8b2d-11f794cb5a3d', 'Economics', 'Books discussing economic theories and concepts', NOW(), NOW()),
    ('f52e1d2c-41a9-4b22-9e94-7c392d8231a4', 'Education', 'Books about teaching methodologies and learning', NOW(), NOW()),
    ('2b1e5c3a-7209-41b2-9e3d-df894eb5a3f5', 'DIY', 'Books providing do-it-yourself projects', NOW(), NOW()),
    ('0a4f3b2c-5e2d-4b22-9e72-3c7926e5a3c4', 'Gardening', 'Books on plant care and gardening techniques', NOW(), NOW()),
    ('3d1e5c3a-32d8-48a2-8b3d-df994eb5a4b6', 'Crafts', 'Books on various craft techniques', NOW(), NOW()),
    ('f72e1c5a-51a9-4b22-8e94-1c392d8231b3', 'Pets', 'Books about pet care and training', NOW(), NOW()),
    ('e91d2c3a-42a8-4b8a-9e94-1d7d892d784c', 'Parenting', 'Books on raising children', NOW(), NOW()),
    ('4f9c2a3e-6b2d-48e2-9b3d-5f794eb5a5d3', 'Legal', 'Books on law and legal studies', NOW(), NOW()),
    ('d52e1a3c-11a9-4b22-8e54-7c392d8231c2', 'Ethics', 'Books on moral philosophy and ethics', NOW(), NOW()),
    ('3a1e5c3a-22d8-48a2-8b3d-df194eb5a6b1', 'Mythology', 'Books about myths and legends', NOW(), NOW()),
    ('5f92c3a2-6b2d-48e2-9b5d-5f794eb5a7c5', 'Adventure', 'Books featuring exciting adventures and quests', NOW(), NOW()),
    ('d82e1c5a-31a9-4b22-9e54-7c392d8231d4', 'Drama', 'Books featuring dramatic storytelling', NOW(), NOW());


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

INSERT INTO books (id, title, year, category_id, author_id, total_copies, language, created_at, updated_at)
VALUES
    ('23cb2f7a-af2b-4279-b825-806703236698', 'The Silent Observer', 2021, 'c8e3f2a1-6209-4d37-9a1d-df994eb5a8f3', 'c3ffbd64-bc2e-46fa-a3a6-b8c605e7e66d', 15, 'English', '2024-01-01 10:30:00', '2024-01-15 12:45:00'),
    ('6092514b-7dcc-40c2-a2d4-8ae8387618d4', 'Digital Fortress', 1998, '0bd2e615-72d8-4b55-82b2-53f47bcd4c31', 'c0e21c1b-b9b8-4ec5-9ed0-d7ccff99377d', 30, 'French', '2023-11-20 08:15:00', '2023-12-05 09:20:00'),
    ('8c2252bf-37e0-4027-a65c-cd4c09f83358', 'The Enchanted Forest', 2015, 'c8e3f2a1-6209-4d37-9a1d-df994eb5a8f3', '6fa500e3-fd12-45a5-a8f6-221c8a5b4d8b', 25, 'Spanish', '2022-06-12 14:00:00', '2022-07-20 16:30:00'),
    ('fca4b6a0-2bfc-494e-8696-7e9c01395871', 'History of the Future', 2020, '0f5c5e24-3e5a-4b55-b2c2-55f45fcd4d12', '0218e409-dc3c-4743-b040-8ea380fa6700', 10, 'German', '2021-09-05 17:45:00', '2021-10-10 19:10:00'),
    ('18fde869-22d5-41d2-b9cf-b98636b41b0e', 'Cooking for Beginners', 2019, '3a1e5c3a-22d8-48a2-8b3d-df194eb5a6b1', 'ed67da24-3da2-4c35-9ae8-64601ba2d843', 40, 'Italian', '2023-03-15 11:10:00', '2023-04-22 13:50:00'),
    ('fb48072f-50ae-4269-9b4a-06cf72ad9dcb', 'Astrophysics Simplified', 2022, '7a5c2e1b-92f8-4c55-82b2-59f45ecd4c88', 'a056f2e7-fb7d-462d-9513-45d80f91f981', 12, 'English', '2022-08-30 15:20:00', '2022-09-15 17:40:00'),
    ('2ce38e8e-ca43-4878-8970-2902d9562636', 'The Last Warrior', 2018, '2bd91e5c-72f8-4d55-92b2-67f45bcd4c99', '6de4c070-0a8d-44d7-8ffb-55a777e0f4f3', 18, 'Portuguese', '2020-12-12 09:45:00', '2021-01-25 11:55:00'),
    ('4ad2d4fb-f209-480f-94db-5cf86f897be9', 'Secrets of the Mind', 2017, '0f5c5e24-3e5a-4b55-b2c2-55f45fcd4d12', 'e0583649-48bb-4f8c-804e-38d5fb7eec5e', 22, 'Russian', '2023-07-07 12:30:00', '2023-08-01 14:15:00'),
    ('e6aed76a-b0a1-4ef8-ac13-657ddfd24a15', 'Travel Beyond Borders', 2023, 'f0cb5e0a-0c95-43c4-81d3-6f97936c477d', '5f0bafc5-066b-4b92-89d1-6a9de75518fd', 35, 'Chinese', '2024-05-05 10:00:00', '2024-06-10 11:30:00'),
    ('a6e0dae2-9fdb-4b7b-8fee-a4be6ff80761', 'The Hidden Code', 2021, '5f92c3a2-6b2d-48e2-9b5d-5f794eb5a7c5', '998c1327-d758-48d7-9b9e-01a3a4e056d0', 28, 'Japanese', '2021-11-18 13:40:00', '2021-12-20 15:00:00'),
    ('fc2a831c-e997-4ec7-884a-e228af55b2b2', 'Tales from the North', 2016, '0f5c5e24-3e5a-4b55-b2c2-55f45fcd4d12', 'db7e6c89-f9ac-4299-8c07-6e1f5e5a67f1', 9, 'English', '2023-05-14 14:23:45', '2024-01-10 10:45:30'),
    ('f12bbee7-2fb2-474c-bd96-60ba8d59ec64', 'Mindful Living', 2020, 'f72e1c5a-51a9-4b22-8e94-1c392d8231b3', 'f20ba955-3304-4432-933c-62b5d17b4bc5', 12, 'French', '2022-07-08 09:34:12', '2023-11-22 16:50:18'),
    ('80c6993d-6368-4633-b554-9c81226ea0b6', 'Artificial Intelligence Today', 2021, 'ef5a6b1a-5e6d-43b2-9e92-2c7d832d745d', '5e6e5930-d9b0-4967-b442-3df7e6984e1c', 7, 'Spanish', '2021-12-01 17:12:29', '2023-09-15 11:20:55'),
    ('127f2114-8c86-469d-af97-b0b09af1353d', 'The Art of Painting', 2015, 'e91d2c3a-42a8-4b8a-9e94-1d7d892d784c', '27a06f17-9351-42c0-8249-58c8da38fd2e', 15, 'Italian', '2020-02-14 08:45:03', '2023-06-30 14:10:21'),
    ('f3a4e1b4-ecc0-477a-85f9-eec218cd8793', 'Legends of the East', 2018, 'f72e1c5a-51a9-4b22-8e94-1c392d8231b3', 'a809d7d3-bb0a-4b72-9379-617b2d2347fc', 20, 'German', '2021-10-10 13:33:44', '2024-01-02 09:05:17'),
    ('07234180-e178-4852-98b9-d4158fe94c9a', 'Mastering Chess', 2017, 'b1e5c3a1-6229-42b2-9e5d-df794eb5a2f3', '8725d975-1956-46c8-b4a5-bf3f914bd8f8', 11, 'Portuguese', '2020-09-25 20:17:56', '2023-04-18 15:30:50'),
    ('9a87e24a-a691-42e9-9345-6457e123b2b7', 'The Hidden Path', 2019, '7cc92a3d-8f39-414f-80c7-49b4c1201bd7', '5124f7a1-c8b9-450d-9b66-c9a2d4d5e9ac', 6, 'Russian', '2023-01-05 18:42:21', '2023-12-10 12:27:48'),
    ('024a961f-e867-45a3-ba3c-941d934595c3', 'Beyond the Stars', 2022, 'd4f212aa-8b3d-46a3-9d3e-3084c2831dc1', '02048d39-6cd6-4940-88f3-64d1bdf88d96', 8, 'Japanese', '2022-04-03 07:11:10', '2024-01-15 19:45:34'),
    ('97fabcbb-7a3a-43f2-b545-97d99b99fab4', 'Gardening 101', 2023, 'f72e1c5a-51a9-4b22-8e94-1c392d8231b3', '3435b61f-bbd1-47cb-9937-ec54d61e54f2', 14, 'Chinese', '2023-06-20 09:55:32', '2023-12-30 17:23:19'),
    ('fe4714bc-3100-4497-98f3-c313ae7b63ad', 'The Lost Civilization', 2016, '5f92c3a2-6b2d-48e2-9b5d-5f794eb5a7c5', '94e0d276-cf16-4ba5-a137-e63d3cfb9db5', 10, 'Arabic', '2021-03-12 11:29:45', '2023-05-09 14:35:22');

-- Insert sample data into `loans`
INSERT INTO loans (id, user_id, book_id, start_date, due_date, return_date, status, created_at, updated_at)
VALUES
    ('dd69a43a-f51f-4f03-9302-c96f13b6b90b', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'fc2a831c-e997-4ec7-884a-e228af55b2b2', '2024-01-01', '2024-01-15', NULL, 'active', NOW(), NOW()),
    ('b63c0269-df60-49cb-b815-685d8a587d6a', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'f12bbee7-2fb2-474c-bd96-60ba8d59ec64', '2024-02-01', '2024-02-15', '2024-02-10', 'completed', NOW(), NOW());

-- Insert sample data into `reservations`
INSERT INTO reservations (id, user_id, book_id, created_at, expiry_date, updated_at)
VALUES
    ('ecdc1e64-8be3-4857-bf2f-c5a54d79f76a', 'a4dfe207-f3bc-4710-8ba4-d9f69a45c093', '80c6993d-6368-4633-b554-9c81226ea0b6', NOW(), '2025-03-01', NOW()),
    ('dd97be09-9467-4b22-91a8-f3c4bb743cc7', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '127f2114-8c86-469d-af97-b0b09af1353d', NOW(), '2025-02-20', NOW());

-- Insert sample data into `ratings`
INSERT INTO ratings (id, user_id, book_id, content, value, created_at, updated_at)
VALUES
    ('6bded01c-0207-42ed-8b09-5455b8167e25', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'f3a4e1b4-ecc0-477a-85f9-eec218cd8793', 'Great book!', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d7139641d', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', '07234180-e178-4852-98b9-d4158fe94c9a', 'Informative but hard to read.', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e21', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '9a87e24a-a691-42e9-9345-6457e123b2b7', '3', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396412', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', '024a961f-e867-45a3-ba3c-941d934595c3', '4', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e23', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '97fabcbb-7a3a-43f2-b545-97d99b99fab4', '5', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396414', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'fe4714bc-3100-4497-98f3-c313ae7b63ad', '6', 3, NOW(), NOW()),

    ('6bded01c-0207-42ed-8b09-5455b8167e26', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'fc2a831c-e997-4ec7-884a-e228af55b2b2', '7', 5, NOW(), NOW()),
    ('8a90d7fc-dc2d-4e96-b6d4-9a3d71396417', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'f12bbee7-2fb2-474c-bd96-60ba8d59ec64', '8', 3, NOW(), NOW());

COMMIT;
