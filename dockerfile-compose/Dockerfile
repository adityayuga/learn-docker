FROM golang:1.14.4

RUN mkdir /my-project
WORKDIR /my-project

COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

RUN go build -o /usr/bin/my-golang

# wait for postgres
RUN chmod +x wait-for-postgres.sh

CMD ["./wait-for-postgres.sh", "my-golang"]