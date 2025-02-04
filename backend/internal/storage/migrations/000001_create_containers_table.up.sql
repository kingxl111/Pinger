
CREATE TABLE containers (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    name VARCHAR(70) NOT NULL,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE pings (
    id SERIAL PRIMARY KEY,
    container_id INT NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    last_success_ping TIMESTAMP NOT NULL,
    CONSTRAINT fk_container
       FOREIGN KEY (container_id)
           REFERENCES containers(id)
           ON DELETE CASCADE
           ON UPDATE CASCADE
);