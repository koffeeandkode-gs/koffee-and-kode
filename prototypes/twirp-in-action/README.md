# Twirp in action: one Go RPC

This is a runnable, deliberately small Twirp service. The contract is [`rpc/greeter/service.proto`](rpc/greeter/service.proto). `service.pb.go` and `service.twirp.go` in the same directory are generated code; edit the `.proto` file and regenerate them instead of editing those Go files.

## Run

From this directory, start the server in one terminal:

```sh
go run ./cmd/server
```

In another terminal, send a normal JSON request:

```sh
curl -sS -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Ada"}' \
  http://localhost:8080/twirp/demo.greeter.Greeter/Greet
```

Expected response: `{"message":"Hello, Ada!"}`.

Or use the generated protobuf client:

```sh
go run ./cmd/client
```

Expected output: `Hello, Ada!`

To see a Twirp error, send an empty name. The server returns HTTP 400 with a JSON body containing `"code":"invalid_argument"`.

```sh
curl -i -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":""}' \
  http://localhost:8080/twirp/demo.greeter.Greeter/Greet
```

## Regenerate

Install [`protoc`](https://protobuf.dev/installation/) first, then install the generator versions used here:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.10
go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.3+incompatible
export PATH="$(go env GOPATH)/bin:$PATH"
```

```sh
protoc -I . \
  --go_out=paths=source_relative:. \
  --twirp_out=paths=source_relative:. \
  rpc/greeter/service.proto
```

Twirp generates a typed server interface and both JSON and protobuf clients. This example implements that interface with one method. It uses Go's `net/http`; there is no separate transport setup.

Run `go test ./...` to check that the generated code, server, and client compile. The `curl` and Go client calls above verify the live behavior.
