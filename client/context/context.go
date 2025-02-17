package context

import (
	"context"
	"errors"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	gojwtgrpcctx "github.com/ralvarezdev/go-jwt/grpc/context"
	"google.golang.org/grpc/metadata"
)

// GetOutgoingCtx returns a context with the raw token
func GetOutgoingCtx(ctx context.Context) (context.Context, error) {
	// Get the raw token from the context
	rawToken, err := gojwtgrpcctx.GetCtxRawToken(ctx)
	if err != nil {
		// Check if the raw token is missing
		if errors.Is(err, gojwtgrpcctx.ErrMissingToken) {
			return context.Background(), nil
		}
		return nil, err
	}

	// Append the raw token to the gRPC context
	grpcCtx := metadata.AppendToOutgoingContext(
		context.Background(),
		gojwtgrpc.AuthorizationMetadataKey,
		rawToken,
	)

	return grpcCtx, nil
}
