# buildmeaplaylist-api
API server handling authentication, database writes, BuildMeAPlaylist backend business logic

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

`IP=0.0.0.0 PORT=8080 DB=YOUR_POSTGRESQL_URL ./buildmeaplaylist-api`

# Bash control
Run background process

`IP=0.0.0.0 PORT=8080 ./buildmeaplaylist-api &`

List running background processes

`jobs`

Kill the third running background process

`kill %3`
