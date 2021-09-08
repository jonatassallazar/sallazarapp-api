FROM golang:1.17.0-alpine3.14 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd

RUN go build -o api ./cmd/

EXPOSE 5000 5000

CMD [ "./api" ]