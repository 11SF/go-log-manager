FROM golang:1.18.3-alpine

COPY . /api/go/log
WORKDIR /api/go/log

RUN apk add build-base
RUN go mod download


EXPOSE 8080

# RUN go run server.go

RUN go build -o /docker-golang-log
CMD [ "/docker-golang-log" ]