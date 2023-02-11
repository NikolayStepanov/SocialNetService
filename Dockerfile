FROM golang:1.20rc3-alpine3.17

WORKDIR /go/src/app

COPY go.mod .
RUN go mod download

COPY . .

RUN apk add postgresql-client

RUN go build -o ./bin/webserver ./cmd/app/main.go

CMD ["./bin/webserver"]
