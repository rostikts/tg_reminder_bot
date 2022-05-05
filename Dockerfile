FROM golang:alpine

WORKDIR /app
ADD . .

RUN go mod download

RUN go build ./bot/main.go
RUN ls ./bot
ENTRYPOINT ./main