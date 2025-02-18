FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

RUN apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /tmp/desafiodev
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]