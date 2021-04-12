FROM golang:1.13-stretch

WORKDIR $GOPATH/src/bitbucket.org/eJaw-all/users-rest-api/

COPY . .

RUN go build -o users-rest-api -v ./cmd/main.go