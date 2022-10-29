# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env .env
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN #go clean --cache
RUN #go build -o main .

#EXPOSE 8081

#CMD [ "./main" ]

CMD "air"