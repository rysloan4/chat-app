CREATE TABLE user (
	uuid VARCHAR(36) NOT NULL,
	username varchar(128) NOT NULL,
	last_seen timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
