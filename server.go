package main

import (
	"github.com/siwik75/grpcexample/internal/api/v1/chat"
	proto "github.com/siwik75/grpcexample/internal/generated/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	proto.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}

}
