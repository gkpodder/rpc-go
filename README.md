# RPC-Go

RPC-Go is a project that demonstrates the use of gRPC in Go. It includes examples of different types of RPCs including Unary, Server Streaming, Client Streaming, and Bidirectional Streaming.

## Features

- Unary RPCs where the client sends a single request to the server and gets a single response back.
- Server streaming RPCs where the client sends a request to the server and gets a stream to read a sequence of messages back.
- Client streaming RPCs where the client writes a sequence of messages and sends them to the server, again using a provided stream.
- Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream.

## Running the Code

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. To run the server, navigate to the `server` directory and run `go run *.go`.
4. To run the client, navigate to the `client` directory and run `go run *.go` in a separate terminal session.

## Dependencies

This project uses the following dependencies:

- [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc)
- [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf)

## Contributing

Contributions are welcome. Please submit a pull request or create an issue to discuss the changes.

## Next Steps?

DOCKERIZE!!!
Try to create a system of microservices and see how the service performs