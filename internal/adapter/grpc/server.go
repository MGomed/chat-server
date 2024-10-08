package grpc_adapter

import (
	"context"
	"log"
	"net"

	domain "github.com/MGomed/chat_server/internal/domain"
	api "github.com/MGomed/chat_server/pkg/chat_api"

	gofakeit "github.com/brianvoe/gofakeit"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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

	logger  *log.Logger
	usecase ChatAPIUsecase
}

// NewGrpcServer is server constructor
func NewGrpcServer(logger *log.Logger, usecase ChatAPIUsecase) *server {
	return &server{
		logger:  logger,
		usecase: usecase,
	}
}

// Serve gets net.Listener and bind it to grpc server,
// also blocking execution by calling Serve()
func (s *server) Serve(listener net.Listener) error {
	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterChatAPIServer(server, s)

	return server.Serve(listener)
}

// Create creates new chat
func (s *server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received create request:\n%s", msg)

	//TODO correct the call ufter implementation
	_, _ = s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))

	resp := &api.CreateResponse{
		Id: gofakeit.Int64(),
	}

	msg, _ = opt.Marshal(resp)
	s.logger.Printf(">>>> Sent create response:\n%s", msg)

	return resp, nil
}

// Delete removes chat by id
func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*emptypb.Empty, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received delete request:\n%s", msg)

	//TODO correct the call ufter implementation
	_ = s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))

	return &emptypb.Empty{}, nil
}

// SendMessage send user's message to chat
func (s *server) SendMessage(ctx context.Context, req *api.SendRequest) (*emptypb.Empty, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received send message request:\n%s", msg)

	//TODO correct the call ufter implementation
	_ = s.usecase.SendMessage(ctx, domain.SendReqFromAPIToDomain(req))

	return &emptypb.Empty{}, nil
}
