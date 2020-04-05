FROM golang:1.12-alpine

RUN apk add --no-cache git

# Create and add to app directory
RUN mkdir /app
ADD . /app/

# Set the Current Working Directory inside the container
WORKDIR /app

# Get dependencies and build the Go app
RUN go mod download
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/main"]
