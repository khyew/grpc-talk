package main

import (
	"errors"
	"fmt"

	pb "github.com/khyew/grpc-talk"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

// Hello is an implementation of the gRPC-defined rpc method
// It responds to clients with an appropriate reply for their hellos
func (h *HelloServer) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	pr, _ := peer.FromContext(ctx)
	name := request.GetName()

	if name == "" {
		fmt.Printf("Received a hello from %v but it contained no name; responding with an error\n", pr.Addr)
		return nil, errors.New("Didn't receive a name")
	}

	fmt.Printf("Received a hello from %v at %v. Sending a kind reply :)\n", name, pr.Addr)
	response := &pb.HelloResponse{
		Reply: fmt.Sprintf("Hello %v! Nice to meet you :)", name),
	}

	return response, nil
}

// HelloServer is an implementation of the demo protobuf-defined server interface
type HelloServer struct{}
