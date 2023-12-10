BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS note(
    id SERIAL,
    author_first_name TEXT,
    aurhor_last_name TEXT,
    note TEXT,
    CONSTRAINT note_pk PRIMARY KEY (id)
);

COMMIT;
