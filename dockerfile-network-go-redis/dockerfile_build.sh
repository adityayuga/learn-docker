# build with image name
docker build --tag app-golang:0.4 .

# create container
docker container create --name golang4 -p 8004:8080 -e REDIS_HOST=redis1 -e REDIS_PORT=6379 app-golang:0.4

# start container
docker container start golang4
docker container logs golang4

# start redis1
docker container start redis1

# create network
docker network create my_golang_network
docker network connect my_golang_network redis1
docker network connect my_golang_network golang4
