package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-grpc/chat"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	
	grpcServer := grpc.NewServer()
	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
