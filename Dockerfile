FROM golang:1.19.1

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY ./backend .
RUN go mod tidy