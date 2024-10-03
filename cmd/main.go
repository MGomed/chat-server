package main

import (
	"fmt"
	"log"
	"net"

	grpc_adapter "github.com/MGomed/chat_server/internal/adapter/grpc"
	chat_api "github.com/MGomed/chat_server/internal/usecase/chat_api"
)

const (
	grpcPort = 50051
)

func main() {
	chatAPIUsecase := chat_api.NewChatAPIUsecase()

	server := grpc_adapter.NewGrpcServer(chatAPIUsecase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
