# build with image name
docker build --tag app-golang:0.3 .

# create container
docker container create --name golang3 -p 8003:8080 -e ENV=production -e NAME=daniel app-golang:0.3

# start container
docker container start golang3
docker container logs golang3