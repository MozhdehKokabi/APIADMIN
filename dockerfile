# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

# RUN go build  -o /go-docker-demo

EXPOSE 3000

CMD [ "go","run","main.go" ]
