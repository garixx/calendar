CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       login character varying(10) NOT NULL UNIQUE,
                       password character varying(10) NOT NULL,
                       username character varying(10) NOT NULL,
                       created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO users (id, login, password, username, created_at) VALUES (1, 'testuser', 'a12345', 'me', '2020-10-05 14:01:10-08');

CREATE TABLE tokens (
                        id SERIAL PRIMARY KEY,
                        login character varying(10) NOT NULL UNIQUE,
                        token character varying(255) NOT NULL UNIQUE,
                        active bool DEFAULT FALSE NOT NULL
);

INSERT INTO tokens (id, login, token, active) VALUES
    (1, 'testuser', 'hfkdshfkjsd', true);