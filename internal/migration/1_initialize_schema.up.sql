CREATE TYPE "user_status_type" AS ENUM ('ACTIVE', 'BLOCKED', 'CLOSED', 'DORMANT');
CREATE TYPE "account_category_type" AS ENUM ('DEBIT', 'CREDIT', 'LOAN');

CREATE TABLE IF NOT EXISTS "role" (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by BIGINT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ,
    deleted_by BIGINT
);

INSERT INTO "role"(id, name, description) VALUES
    (1, 'Super Admin', 'All access are granted'),
    (2, 'Admin', 'Several access are granted'),
    (3, 'User', null)
ON CONFLICT ON CONSTRAINT role_pkey DO NOTHING;

CREATE TABLE IF NOT EXISTS "user" (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    role_id BIGINT NOT NULL REFERENCES "role"(id),
    password BYTEA NOT NULL,
    passcode BYTEA NOT NULL,
    status user_status_type NOT NULL,
    is_email_verified BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by BIGINT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ,
    deleted_by BIGINT
);

CREATE TABLE IF NOT EXISTS "account" (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    user_id BIGINT NOT NULL REFERENCES "user"(id),
    category account_category_type NOT NULL,
    balance NUMERIC(16, 2) NOT NULL DEFAULT '0.00',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by BIGINT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ,
    deleted_by BIGINT
);

CREATE TABLE IF NOT EXISTS "account_history" (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    account_id BIGINT NOT NULL REFERENCES "account"(id),
    amount NUMERIC(16, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by BIGINT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ,
    deleted_by BIGINT
);

CREATE UNIQUE INDEX user_username_idx ON "user" (LOWER(username));
CREATE UNIQUE INDEX user_email_idx ON "user" (LOWER(email));
CREATE UNIQUE INDEX user_phone_number_idx ON "user" (LOWER(phone_number));
