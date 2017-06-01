/// Server component of grpc-in-docker to explore functionality in scaled docker containers
/// Author: Chris Clarkson
/// Date: 1st June 2017
/// Company: CJC Software Solutions ltd.

package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/ClarksonCJ/grpc-in-docker/pkg/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.String("Port", ":30551", "Port number to listen for requests on")

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Greet Request: %s", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
