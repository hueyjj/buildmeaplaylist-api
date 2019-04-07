# buildmeaplaylist-api
API server handling authentication, database writes, BuildMeAPlaylist backend business logic

# Requirements
Install buildmeaplaylist-api

`go get -u github.com/hueyjj/buildmeaplaylist-api`

Install request router

`go get -u github.com/gorilla/mux`

# Running
Start the server

`IP=0.0.0.0 PORT=8080 ./buildmeaplaylist-api`

# Bash control
Run background process

`IP=0.0.0.0 PORT=8080 ./buildmeaplaylist-api &`

List running background processes

`jobs`

Kill the third running background process

`kill %3`
