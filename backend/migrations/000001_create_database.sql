CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INTEGER NOT NULL
);

CREATE TABLE avatars (
    name VARCHAR(255) PRIMARY KEY,
    image_url VARCHAR(255) NOT NULL,
    interests TEXT[]
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    is_completed BOOLEAN
);
