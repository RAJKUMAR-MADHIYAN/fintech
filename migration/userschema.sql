
CREATE TYPE user_type AS ENUM ('admin', 'regular', 'guest');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    user_type user_type NOT NULL
);