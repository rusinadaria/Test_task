CREATE TABLE songs (
    id  SERIAL PRIMARY KEY,
    song VARCHAR(255) NOT NULL UNIQUE,
    group_name VARCHAR(255),
    release_date VARCHAR(255),
    text TEXT,
    link VARCHAR(255)
)