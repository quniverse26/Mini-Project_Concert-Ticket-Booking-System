# syntax=docker/dockerfile:experimental

# specify the base image to be used for applicatin, alpine or ubuntu
FROM golang:1.17-alpine

# specify a qorking directory inside the image
WORKDIR /app

# create Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile application
RUN go build -o /github.com/quniverse26/miniproject

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080