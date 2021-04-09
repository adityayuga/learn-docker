# build with image name
docker build --tag app-golang:0.5 .

# create container
docker container create --name golang5 app-golang:0.5