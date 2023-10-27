CREATE TABLE IF NOT EXISTS user_info (
    id VARCHAR PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL DEFAULT '',
    phone_number VARCHAR(256) NOT NULL DEFAULT '',
    photo_url VARCHAR(512) NOT NULL,
    disabled BOOLEAN NOT NULL,
    email_verified BOOLEAN NOT NULL,
    dob VARCHAR(256) NOT NULL DEFAULT '',
    hash_tag VARCHAR(256) NOT NULL DEFAULT '',
    pronoun VARCHAR(256) NOT NULL DEFAULT '',
    bio VARCHAR(1024) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

ALTER TABLE
    user_info
ADD
    CONSTRAINT idx_user_info_id_unique_constraint UNIQUE (id);