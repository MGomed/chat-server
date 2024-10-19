package grpc_adapter

import (
	"context"
	"log"
	"net"

	domain "github.com/MGomed/chat_server/internal/domain"
	api "github.com/MGomed/chat_server/pkg/chat_api"

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

// Config declair grpc config interface
type Config interface {
	Address() string
}

type server struct {
	api.UnimplementedChatAPIServer

	logger  *log.Logger
	config  Config
	usecase ChatAPIUsecase
}

// NewGrpcServer is server constructor
func NewGrpcServer(logger *log.Logger, config Config, usecase ChatAPIUsecase) *server {
	return &server{
		logger:  logger,
		config:  config,
		usecase: usecase,
	}
}

// Serve gets net.Listener and bind it to grpc server,
// also blocking execution by calling Serve()
func (s *server) Serve() error {
	lis, err := net.Listen("tcp", s.config.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterChatAPIServer(server, s)

	return server.Serve(lis)
}

// Create creates new chat
func (s *server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	resp, err := s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return domain.CreateRespToAPIFromDomain(resp), nil
}

// Delete removes chat by id
func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*emptypb.Empty, error) {
	err := s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SendMessage send user's message to chat
func (s *server) SendMessage(ctx context.Context, req *api.SendRequest) (*emptypb.Empty, error) {
	err := s.usecase.SendMessage(ctx, domain.SendReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
