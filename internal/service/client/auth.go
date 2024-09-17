package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationMetadataKey = "authorization"
)

func (c *ClientImpl) AuthInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.Pairs(AuthorizationMetadataKey, c.authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
