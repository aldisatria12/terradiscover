CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	username VARCHAR(18),
	password VARCHAR(100),
	email VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS contacts (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL REFERENCES users (id),
	name VARCHAR(100),
	phone VARCHAR(15),
	email VARCHAR(100),
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);