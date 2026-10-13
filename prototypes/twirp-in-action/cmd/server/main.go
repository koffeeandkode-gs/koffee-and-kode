package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/gautham-surakanti/koffee-and-kode/prototypes/twirp-in-action/rpc/greeter"
	"github.com/twitchtv/twirp"
)

type greeter struct{}

func (greeter) Greet(_ context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	if req.Name == "" {
		return nil, twirp.InvalidArgumentError("name", "must not be empty")
	}
	return &pb.GreetResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	log.Println("Twirp server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", pb.NewGreeterServer(greeter{})))
}
