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

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    dep TEXT NOT NULL,
    lvl TEXT NOT NULL,
    course INT NOT NULL,
    sport TEXT,
    club TEXT,
    mer TEXT,
    FOREIGN KEY (avatar_id) REFERENCES avatars(id)
);

CREATE TABLE users_achievements (
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (achievement_id) REFERENCES achievements(id)
);

CREATE TABLE users_tasks (
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);