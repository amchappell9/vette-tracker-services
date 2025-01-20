CREATE TABLE IF NOT EXISTS vettes (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL,
    miles VARCHAR(255) NOT NULL,
    cost VARCHAR(255) NOT NULL,
    transmission_type VARCHAR(255) NOT NULL,
    exterior_color VARCHAR(255) NOT NULL,
    interior_color VARCHAR(255) NOT NULL,
    submodel VARCHAR(255) NOT NULL,
    trim VARCHAR(255) NOT NULL,
    packages TEXT[] NOT NULL,
    link TEXT NOT NULL
);