package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/chat_server/consts"
	errors "github.com/MGomed/chat_server/internal/config/errors"
)

type accessConfig struct {
	host string
	port string
}

// NewAccessConfig is grpcConfig struct constructor
func NewAccessConfig() (*accessConfig, error) {
	host := os.Getenv(consts.AccessServiceHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", errors.ErrEnvNotFound, consts.AccessServiceHostEnv)
	}

	port := os.Getenv(consts.AccessServicePortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", errors.ErrEnvNotFound, consts.AccessServicePortEnv)
	}

	return &accessConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns grpc ip address
func (c *accessConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
