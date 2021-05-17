FROM golang:1.13-stretch

WORKDIR $GOPATH/src/bitbucket.org/electronicjaw/novels-backend/

COPY . .

RUN go build -o novels-backend -v ./cmd/main.go