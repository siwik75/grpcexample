package main

import (
	"context"
	chat "github.com/siwik75/grpcexample/internal/generated/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello %v", err)
	}
	log.Printf("response from Server : %s", response.Body)
}
