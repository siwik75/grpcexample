package chat

import (
	proto "github.com/siwik75/grpcexample/internal/generated/api"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *proto.Message) (*proto.Message, error) {
	log.Printf("Received Message body from client: %s", message.Body)
	return &proto.Message{Body: "Hello from the Server!"}, nil
}
