CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INTEGER NOT NULL
);

CREATE TABLE avatars (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    image_url VARCHAR(255) NOT NULL,
    interests TEXT[]
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    is_completed BOOLEAN
);
