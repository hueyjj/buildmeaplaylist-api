# Postgresql
Not sure If we want to automate creating and validating tables yet. But when we do, we definitely need to make sure to check if tables exist (not malformed) or don't exist. 

Current tables we need (manually make these for now):

`CREATE TABLE members(email varchar(50) NOT NULL PRIMARY KEY, password varchar (100) NOT NULL, first_name varchar(20) NOT NULL, last_name varchar(20) NOT NULL);`

## Database url 
The format looks something like this.

`postgres://user:password@localhost:5432/buildmeaplaylist`

## Running

Login psql and connect to database

`sudo -u USER psql -d DATABASE`

`sudo -u jasperjeng psql -d buildmeaplaylist`

## psql commands
Quit

`\q`

List tables, views, and sequences

`\d`

Describe table

`\d members`

List rows

`table members;`

Drop table

`drop table members;`