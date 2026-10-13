package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/gautham-surakanti/koffee-and-kode/prototypes/twirp-in-action/rpc/greeter"
)

func main() {
	client := pb.NewGreeterProtobufClient("http://localhost:8080", &http.Client{})
	response, err := client.Greet(context.Background(), &pb.GreetRequest{Name: "Ada"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Message)
}
