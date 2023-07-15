
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	uuid VARCHAR(255) NOT NULL UNIQUE,
	name VARCHAR(255),
	email VARCHAR(255),
	password VARCHAR(255),
	created_at DATETIME,
	updated_at DATETIME
);

-- +migrate Down
DROP TABLE IF EXISTS user;
