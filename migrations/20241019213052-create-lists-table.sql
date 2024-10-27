-- +migrate Up
CREATE TABLE lists (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING NOT NULL,
    uuid CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX lists_uuid_index ON lists (uuid);

-- +migrate Down
DROP INDEX lists_uuid_index;

DROP TABLE lists;
