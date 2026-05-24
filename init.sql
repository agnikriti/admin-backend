CREATE TABLE email_requests (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    email VARCHAR(255),
    mobile VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW()
);