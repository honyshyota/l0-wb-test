FROM golang:latest

RUN go version

COPY . /go/src/app

WORKDIR /go/src/app

# install psql
RUN apt-get update

RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable

RUN chmod +x wait-for-postgres.sh

# build go app

RUN go mod download

RUN go build -o api cmd/main.go

EXPOSE 8080

CMD ["./api"]