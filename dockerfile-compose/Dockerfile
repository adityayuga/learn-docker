FROM golang:1.14.4

RUN mkdir /my-project
WORKDIR /my-project

COPY go.mod .
COPY go.sum .

COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

RUN go build -o /usr/bin/my-golang

CMD ["my-golang"]