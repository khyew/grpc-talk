package main

import (
	"fmt"
	"net"

	pb "github.com/khyew/grpc-talk"
	"google.golang.org/grpc"
)

const HOST = "0.0.0.0"
const PORT = "6666"

func main() {
	ln, err := net.Listen("tcp", net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Printf("Couldn't bind to port %v because: %v\n", PORT, err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDemoServiceServer(grpcServer, &HelloServer{})

	fmt.Printf("gRPC server now listening on port %v\n", PORT)
	grpcServer.Serve(ln)
}
