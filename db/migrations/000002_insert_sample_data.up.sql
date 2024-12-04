BEGIN;

-- Insert sample data into `users`
-- INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
-- VALUES
-- ('548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'User One', 'user1@example.com', 'passwordhash1', NOW(), NOW()),
-- ('c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'User Two', 'user2@example.com', 'passwordhash2', NOW(), NOW()),
-- ('a4dfe207-f3bc-4710-8ba4-d9f69a45c093', 'User Three', 'user3@example.com', 'passwordhash3', NOW(), NOW());

-- Insert sample data into `categories`
INSERT INTO categories (id, name, description, created_at, updated_at)
VALUES
('08d4b7cf-5618-42fa-943b-854e00e65d22', 'Fiction', 'Category for fictional books', NOW(), NOW()),
('7cc92a3d-8f39-414f-80c7-49b4c1201bd7', 'Science', 'Category for science books', NOW(), NOW());

-- Insert sample data into `authors`
INSERT INTO authors (id, first_name, last_name, nationality, birth_date, death_date, bio, website, created_at, updated_at)
VALUES
('e2435d88-dc2d-4c18-9068-85c3b5d885cd', 'Author', 'One', 'American', '1980-01-01', NULL, 'Bio of Author One', 'https://authorone.com', NOW(), NOW()),
('f9184633-798d-4b4f-9c4b-e01432683418', 'Author', 'Two', 'British', '1975-06-15', NULL, 'Bio of Author Two', NULL, NOW(), NOW());

-- Insert sample data into `books`
INSERT INTO books (id, title, year, author_id, category_id, total_copies, language, created_at, updated_at)
VALUES
('27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', 'Book One', 2020, 'e2435d88-dc2d-4c18-9068-85c3b5d885cd', '08d4b7cf-5618-42fa-943b-854e00e65d22', 5, 'English', NOW(), NOW()),
('bc582d8b-cbeb-45e6-8b0b-c89e871b5833', 'Book Two', 2018, 'f9184633-798d-4b4f-9c4b-e01432683418', '7cc92a3d-8f39-414f-80c7-49b4c1201bd7', 3, 'French', NOW(), NOW()),
('a1c61255-3547-4c6e-b3ba-f0ddf2763791', 'Book Three', 2022, 'e2435d88-dc2d-4c18-9068-85c3b5d885cd', NULL, 2, 'Spanish', NOW(), NOW());

-- Insert sample data into `loans`
-- INSERT INTO loans (id, user_id, book_id, start_date, due_date, return_date, status, created_at, updated_at)
-- VALUES
-- ('dd69a43a-f51f-4f03-9302-c96f13b6b90b', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', '2024-01-01', '2024-01-15', NULL, 'active', NOW(), NOW()),
-- ('b63c0269-df60-49cb-b815-685d8a587d6a', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', '2024-02-01', '2024-02-15', '2024-02-10', 'completed', NOW(), NOW());

-- Insert sample data into `reservations`
-- INSERT INTO reservations (id, user_id, book_id, created_at, expiry_date, updated_at)
-- VALUES
-- ('ecdc1e64-8be3-4857-bf2f-c5a54d79f76a', 'a4dfe207-f3bc-4710-8ba4-d9f69a45c093', 'a1c61255-3547-4c6e-b3ba-f0ddf2763791', NOW(), '2024-03-01', NOW()),
-- ('dd97be09-9467-4b22-91a8-f3c4bb743cc7', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', NOW(), '2024-02-20', NOW());
--
-- -- Insert sample data into `reviews`
-- INSERT INTO reviews (id, user_id, book_id, content, rating, created_at, updated_at)
-- VALUES
-- ('6bded01c-0207-42ed-8b09-5455b8167e25', '548f3b6c-63a1-4c2d-9e85-6a9f76bc3a90', '27dc7f06-6a2f-41f5-a69f-ec2303bcae3b', 'Great book!', 5, NOW(), NOW()),
-- ('8a90d7fc-dc2d-4e96-b6d4-9a3d7139641d', 'c87bfa5f-9d84-450d-b96e-0e42b8e72673', 'bc582d8b-cbeb-45e6-8b0b-c89e871b5833', 'Informative but hard to read.', 3, NOW(), NOW());
--
COMMIT;
