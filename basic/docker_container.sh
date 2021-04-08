# list command
docker container -h

# create container (without expose port)
docker container create --name redis1 redis:latest
docker container create --name redis2 redis:latest

# start container 
docker container start redis1
docker container start redis2
docker container ls

# stop container
docker container stop redis1 redis2

# delete container (will not delete image)
docker container rm redis1 redis2
docker ps -a

# create container (with expose port)
docker container create --name redis1 -p 6379:6379 redis:latest
docker container create --name redis2 -p 6380:6379 redis:latest

# test redis-cli
redis-cli -p 6379
redis-cli -p 6380
