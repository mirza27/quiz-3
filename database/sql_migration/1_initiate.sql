-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE category(
    id BIGINT NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE book(
    id BIGINT NOT NULL,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year INT NOT NULL,
    price VARCHAR(255),
    total_page INT NOT NULL,
    thickness VARCHAR(256),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    category_id BIGINT NOT NULL
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON book
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


-- +migrate StatementEnd
