-- +migrate Up
CREATE TABLE lists (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING NOT NULL,
    guid CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX lists_guid_index ON lists (guid);

-- +migrate Down
DROP TABLE lists;

DROP INDEX lists_guid_index;
