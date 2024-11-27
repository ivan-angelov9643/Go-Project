BEGIN;

-- Insert sample data into `users`
INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', 'User One', 'user1@example.com', 'passwordhash1', NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', 'User Two', 'user2@example.com', 'passwordhash2', NOW(), NOW()),
('33333333-3333-3333-3333-333333333333', 'User Three', 'user3@example.com', 'passwordhash3', NOW(), NOW());

-- Insert sample data into `categories`
INSERT INTO categories (id, name, description, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', 'Fiction', 'Category for fictional books', NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', 'Science', 'Category for science books', NOW(), NOW());

-- Insert sample data into `authors`
INSERT INTO authors (id, first_name, last_name, nationality, birth_date, death_date, bio, website, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', 'Author', 'One', 'American', '1980-01-01', NULL, 'Bio of Author One', 'https://authorone.com', NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', 'Author', 'Two', 'British', '1975-06-15', NULL, 'Bio of Author Two', NULL, NOW(), NOW());

-- Insert sample data into `books`
INSERT INTO books (id, title, year, author_id, category_id, total_copies, language, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', 'Book One', 2020, '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 5, 'English', NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', 'Book Two', 2018, '22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 3, 'French', NOW(), NOW()),
('33333333-3333-3333-3333-333333333333', 'Book Three', 2022, '11111111-1111-1111-1111-111111111111', NULL, 2, 'Spanish', NOW(), NOW());

-- Insert sample data into `loans`
INSERT INTO loans (id, user_id, book_id, start_date, due_date, return_date, status, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '2024-01-01', '2024-01-15', NULL, 'active', NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', '2024-02-01', '2024-02-15', '2024-02-10', 'completed', NOW(), NOW());

-- Insert sample data into `reservations`
INSERT INTO reservations (id, user_id, book_id, created_at, expiry_date, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', '33333333-3333-3333-3333-333333333333', '33333333-3333-3333-3333-333333333333', NOW(), '2024-03-01', NOW()),
('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', NOW(), '2024-02-20', NOW());

-- Insert sample data into `reviews`
INSERT INTO reviews (id, user_id, book_id, content, rating, created_at, updated_at)
VALUES
('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'Great book!', 5, NOW(), NOW()),
('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222', 'Informative but hard to read.', 3, NOW(), NOW());

COMMIT;

