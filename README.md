# User Managment System

## Create DB
docker-compose up

## Run
First enter DockerFile and comment ```ENTRYPOINT go run server.go --deploy``` and uncomment ```ENTRYPOINT go run server.go```
after that run docker-compose up --build
