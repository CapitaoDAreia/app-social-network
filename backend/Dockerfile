FROM golang:1.19.1 AS build-backend
WORKDIR /usr/src/app-backend

RUN go install github.com/cosmtrek/air@latest 
COPY . .
RUN go mod tidy

