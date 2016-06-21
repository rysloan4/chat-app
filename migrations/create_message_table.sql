CREATE TABLE message (
  	uuid VARCHAR(36) NOT NULL,
  	from_username VARCHAR(128) NOT NULL,
  	to_username VARCHAR(128) NOT NULL,
  	content TEXT,
  	created_at DATETIME NOT NULL
  );