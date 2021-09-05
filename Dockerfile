FROM golang:1.17.0-alpine3.14 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY src ./src

RUN go build ./src/cmd/

EXPOSE 5000 5000

CMD [ "./cmd" ]