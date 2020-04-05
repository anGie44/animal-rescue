# animal-rescue

To run the server on your system:

1. Run `go build -o main` to create the binary (`main`)
2. Run the binary : `./main`

To run the server in a Docker container:

1. Run `docker build -t animal-rescue .`
2. Run `docker run -p $PORT:8080 animal-rescue`. Where `$PORT` is your host's port.  

To run tests:

1. Run `go test ./...`
