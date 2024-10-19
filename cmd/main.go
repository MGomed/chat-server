package main

import (
	"context"
	"flag"
	"log"

	grpc_adapter "github.com/MGomed/chat_server/internal/adapter/grpc"
	postgres "github.com/MGomed/chat_server/internal/adapter/postgres"
	config "github.com/MGomed/chat_server/internal/config"
	env_config "github.com/MGomed/chat_server/internal/config/env"
	chat_api "github.com/MGomed/chat_server/internal/usecase/chat_api"
	logger "github.com/MGomed/chat_server/pkg/logger"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "build/.env", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	err := env_config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := initConfig()

	log, err := logger.InitLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pgAdapter, err := postgres.NewAdapter(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	userAPIUsecase := chat_api.NewChatAPIUsecase(log, pgAdapter)

	server := grpc_adapter.NewGrpcServer(log, cfg, userAPIUsecase)

	log.Println("Starting GRPC server!")

	if err := server.Serve(); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}

func initConfig() *config.Config {
	grpcConfig, err := env_config.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgConfig, err := env_config.NewPgConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerConfig, err := env_config.NewLoggerConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config.NewConfig(grpcConfig, pgConfig, loggerConfig)
}
