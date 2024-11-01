package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/chat_server/consts"
	errors "github.com/MGomed/chat_server/internal/config/errors"
)

type grpcConfig struct {
	host string
	port string
}

// NewAPIConfig is grpcConfig struct constructor
func NewAPIConfig() (*grpcConfig, error) {
	host := os.Getenv(consts.ServerHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", errors.ErrEnvNotFound, consts.ServerHostEnv)
	}

	port := os.Getenv(consts.ServerPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", errors.ErrEnvNotFound, consts.ServerPortEnv)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns grpc ip address
func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
