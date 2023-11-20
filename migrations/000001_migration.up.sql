BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS note(
    id uuid DEFAULT uuid_generate_v4(),
    author_first_name VARCHAR,
    aurhor_last_name VARCHAR,
    note TEXT,
    CONSTRAINT note_pk PRIMARY KEY (id)
);

COMMIT;
