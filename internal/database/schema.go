package database

var Schema = `
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    first_name VARCHAR(50),
	last_name VARCHAR(50),
    email VARCHAR(50),
	created_on TIMESTAMP,
	updated_on TIMESTAMP,
	deleted_on TIMESTAMP
);
`
