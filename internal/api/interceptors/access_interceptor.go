package interceptors

import (
	"context"
	"strings"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	metadata "google.golang.org/grpc/metadata"

	consts "github.com/MGomed/chat_server/consts"
	api_errors "github.com/MGomed/chat_server/internal/api/errors"
	access_api "github.com/MGomed/common/access_api"
)

// AccessInterceptor is interceptor to get access through auth service
func AccessInterceptor(
	accessServAddress string,
) func(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, api_errors.ErrMetadataNotProvided
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, api_errors.ErrHeaderNotProvided
		}

		if !strings.HasPrefix(authHeader[0], consts.AccessPrefix) {
			return nil, api_errors.ErrHeaderWrongFormat
		}

		clientCtx := context.Background()
		clientCtx = metadata.NewOutgoingContext(clientCtx, md)

		conn, err := grpc.NewClient(
			accessServAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, err
		}

		client := access_api.NewAccessAPIClient(conn)

		_, err = client.Check(clientCtx, &access_api.CheckRequest{
			EndpointAddress: info.FullMethod,
		})
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
