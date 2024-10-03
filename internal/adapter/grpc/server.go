package grpc_adapter

import (
	"context"
	"net"

	domain "github.com/MGomed/chat_server/internal/domain"
	api "github.com/MGomed/chat_server/pkg/chat_api"

	gofakeit "github.com/brianvoe/gofakeit"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// ChatAPIUsecase interface of chat_api usecase
type ChatAPIUsecase interface {
	Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error)
	Delete(ctx context.Context, req *domain.DeleteRequest) error
	SendMessage(ctx context.Context, req *domain.SendRequest) error
}

type server struct {
	api.UnimplementedChatAPIServer

	usecase ChatAPIUsecase
}

func NewGrpcServer(usecase ChatAPIUsecase) *server {
	return &server{
		usecase: usecase,
	}
}

func (s *server) Serve(listener net.Listener) error {
	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterChatAPIServer(server, s)

	return server.Serve(listener)
}

func (s *server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	//TODO correct the call ufter implementation
	_, _ = s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))

	return &api.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*emptypb.Empty, error) {
	//TODO correct the call ufter implementation
	_ = s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *api.SendRequest) (*emptypb.Empty, error) {
	//TODO correct the call ufter implementation
	_ = s.usecase.SendMessage(ctx, domain.SendReqFromAPIToDomain(req))

	return &emptypb.Empty{}, nil
}
