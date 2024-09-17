package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophertask/internal/auth"
)

type ContextKey string

const (
	AuthorizationMetadataKey                = "authorization" // AuthorizationMetadataKey is the key used to retrieve the authorization token from gRPC metadata.
	AuthorizationTokenContextKey ContextKey = "user"          // AuthorizationTokenContextKey is the context key used to store the authentication token claims in the context.
)

func (s *Server) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	exemptedMethods := map[string]bool{
		"/gophertask.gophertaskapi.v1.AuthServiceV1/SignUp": true,
		"/gophertask.gophertaskapi.v1.AuthServiceV1/SignIn": true,
	}

	if _, exempt := exemptedMethods[info.FullMethod]; exempt {
		return handler(ctx, req)
	}

	var tokenStr string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Error("Auth interceptor: missing request metadata")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	values := md.Get(AuthorizationMetadataKey)
	if len(values) == 0 {
		s.Logger.Error("Auth interceptor: missing request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}
	tokenStr = values[0]

	if len(tokenStr) == 0 {
		s.Logger.Error("Auth interceptor: missing request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	token, err := auth.ValidateToken(tokenStr, s.Config.AuthenticationKey)
	if err != nil {
		s.Logger.Error("Auth interceptor: invalid request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token")
	}

	ctx = context.WithValue(ctx, AuthorizationTokenContextKey, token.Claims)
	return handler(ctx, req)
}
