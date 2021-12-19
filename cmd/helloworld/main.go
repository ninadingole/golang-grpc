package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ninadingole/golang-grpc/pkg/server/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &helloworld.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
