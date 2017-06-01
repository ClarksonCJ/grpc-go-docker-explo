/// Client application to make requests via gRPC to a dockerised server
/// Author: Chris Clarkson
/// Date: 1st June 2017
/// Company: CJC Software Solutions ltd.
package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/ClarksonCJ/grpc-in-docker/pkg/protobuf"
	"google.golang.org/grpc"
)

var address = flag.String("Address", "localhost:30551", "Server address of the gRPC remote application host")
var name = flag.String("Name", "world", "Default string to send to gRPC server")

func main() {
	flag.Parse()

	// Setup connection to the server
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response
	name := *name
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
