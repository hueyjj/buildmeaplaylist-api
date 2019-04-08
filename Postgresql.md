# Postgresql
Not sure If we want to automate creating and validating tables yet. But when we do, we definitely need to make sure to check if tables exist (not malformed) or don't exist. 

Current tables we need (manually make these for now):

`CREATE TABLE members(email varchar(50) NOT NULL PRIMARY KEY, first_name varchar(20) NOT NULL, last_name varchar(20) NOT NULL);`