# buildmeaplaylist-api
API server handling authentication, database writes, BuildMeAPlaylist backend business logic

# Quickstart
```
go get -u github.com/hueyjj/buildmeaplaylist-api
go get -u ./... # Install all dependencies
go build
IP=0.0.0.0 PORT=8080 DB=YOUR_POSTGRESQL_URL SESSION_KEY=YOUR_SESSION_KEY ~/go/bin/buildmeaplaylist-api

# TODO: How to quick start a database? Fill this in later Jasper
```

# Requirements
Install buildmeaplaylist-api

`go get -u github.com/hueyjj/buildmeaplaylist-api`

Install request router

`go get -u github.com/gorilla/mux`

Install sessions

`go get -u github.com/gorilla/sessions`

Install postgresql driver

`go get -u github.com/lib/pq`

Install bcrypt for hashing

`go get -u golang.org/x/crypto/bcrypt`

# Running
Start the server

`IP=0.0.0.0 PORT=8080 DB=YOUR_POSTGRESQL_URL SESSION_KEY=YOUR_SESSION_KEY ./buildmeaplaylist-api`

# Bash control
Run background process

`PROGRAM &`

List running background processes

`jobs`

Kill the third running background process

`kill %3`

# Testing
Tool: Postman
Routes:
    localhost:8080/api/signup
    localhost:8080/api/login
Body: x-www-form-urlencoded

signup values: email, firstName, lastName, password
login values: email, password
