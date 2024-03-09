SET timezone = 'America/Sao_Paulo';
CREATE SCHEMA IF NOT EXISTS rinha;

CREATE TABLE IF NOT EXISTS rinha.users
(
    id              SERIAL PRIMARY KEY,
    limit_in_cents  INTEGER NOT NULL,
    initial_balance INTEGER NOT NULL DEFAULT 0
);

INSERT INTO rinha.users (id, limit_in_cents, initial_balance)
VALUES (DEFAULT, 1000 * 100, 0),
       (DEFAULT, 800 * 100, 0),
       (DEFAULT, 10000 * 100, 0),
       (DEFAULT, 100000 * 100, 0),
       (DEFAULT, 5000 * 100, 0);

CREATE TABLE IF NOT EXISTS rinha.history
(
    id          SERIAL PRIMARY KEY,
    user_id     SMALLINT    NOT NULL,
    value       INTEGER     NOT NULL,
    type        CHAR(1)     NOT NULL,
    description VARCHAR(10) NOT NULL,
    do_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    rinha.history
    SET
        (autovacuum_enabled = false);

CREATE INDEX idx_history ON rinha.history (user_id);
