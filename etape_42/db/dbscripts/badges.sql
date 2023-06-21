/* Create Tables */

CREATE DATABASE badges_db OWNER postgres;

/* CONNECT TO badges_db USER postgres USING postgres; */
\c badges_db;

CREATE TABLE users
(
	id int NOT NULL,
	
	PRIMARY KEY (id)
	
) WITHOUT OIDS;

CREATE TABLE badges
(
	id int NOT NULL,

	name text NOT NULL,
	url text NOT NULL,

	user_id int NOT NULL,
	
	PRIMARY KEY (id),
	CONSTRAINT fk_user
	  FOREIGN KEY (user_id) 
	    REFERENCES users(id)
	
) WITHOUT OIDS;
