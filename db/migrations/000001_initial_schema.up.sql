BEGIN;

CREATE TABLE category (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(5000),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE author (
    id UUID PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    death_date TIMESTAMP,
    bio VARCHAR(5000),
    website VARCHAR(500),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE book (
    id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    year VARCHAR(4) NOT NULL,
    author_id UUID REFERENCES author(id) ON DELETE CASCADE,
    category_id UUID REFERENCES category(id) ON DELETE SET NULL,
    total_copies INT DEFAULT 1,
    language VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE loan (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES user(id) ON DELETE CASCADE,
    book_id UUID REFERENCES book(id) ON DELETE CASCADE,
    start_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    status ENUM('active', 'completed') NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE reservation (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES user(id) ON DELETE CASCADE,
    book_id UUID REFERENCES book(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    expiry_date TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE review (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES user(id) ON DELETE CASCADE,
    book_id UUID REFERENCES book(id) ON DELETE CASCADE,
    content VARCHAR(5000) NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

COMMIT;