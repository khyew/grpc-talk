package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/khyew/grpc-talk"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %v IP PORT NAME\n", os.Args[0])
		return
	}
	host := os.Args[1]
	port := os.Args[2]
	name := os.Args[3]

	conn, err := grpc.Dial(net.JoinHostPort(host, port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to server at (%v:%v):%v\n", host, port, err)
		return
	}
	defer conn.Close()

	demoClient := pb.NewDemoServiceClient(conn)
	request := pb.HelloRequest{
		Name: name,
	}
	response, err := demoClient.Hello(context.Background(), &request)
	if err != nil {
		fmt.Printf("Error calling remote procedure 'Hello': %v\n", err)
		return
	}

	fmt.Printf("%v\n", response.Reply)
}
