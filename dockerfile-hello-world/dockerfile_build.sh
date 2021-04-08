# build with image name
docker build --tag app-golang:0.1 .

# create container
docker container create --name golang1 -p 8001:8080 app-golang:0.1

# start container
docker container start golang1
docker container logs golang1

# build with image name with custom dockerfilename
docker build --tag app-golang:0.2 -f Dockerfile-test .

# create container
docker container create --name golang2 -p 8002:8080 app-golang:0.2

# login to golang2
docker exec -t -i golang2 /bin/bash