# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="muhammad huzair <muhammadhuzair@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR $GOPATH/src/go-echo-example

# Copy the source from the current directory to the working Directory inside the container
#COPY dockerfiles .
COPY . .

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o go-echo-example

# Expose port 8080 to the outside world
EXPOSE 3010

#Command to run the executable
CMD ["./go-echo-example"]