CREATE TABLE containers (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    last_success_ping TIMESTAMP NOT NULL
);