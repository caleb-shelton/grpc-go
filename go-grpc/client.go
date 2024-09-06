package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"go-grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHelloTwo(context.Background(), &chat.NumberMessage{Body: 100})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %d", response.Number)
}
