package server

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophertask/internal/auth"
	config "github.com/Stern-Ritter/gophertask/internal/config/server"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

func TestAuthInterceptor(t *testing.T) {
	secret := "secret"
	validAuthToken, err := auth.NewToken(model.User{ID: "1", Login: "login", Password: "password"}, secret, time.Hour)
	require.NoError(t, err, "unexpected error creating valid auth token")
	invalidAuthToken, err := auth.NewToken(model.User{ID: "1", Login: "login", Password: "password"}, "invalid secret", time.Hour)
	require.NoError(t, err, "unexpected error creating valid auth token")

	tests := []struct {
		name           string
		authToken      string
		exempted       bool
		expectedStatus codes.Code
	}{
		{
			name:           "should allow request when token is valid and method is not exempted",
			authToken:      validAuthToken,
			exempted:       false,
			expectedStatus: codes.OK,
		},
		{
			name:           "should allow request for exempted methods",
			authToken:      "",
			exempted:       true,
			expectedStatus: codes.OK,
		},
		{
			name:           "should reject request when token is missing",
			authToken:      "",
			exempted:       false,
			expectedStatus: codes.Unauthenticated,
		},
		{
			name:           "should reject request when token is invalid",
			authToken:      invalidAuthToken,
			exempted:       false,
			expectedStatus: codes.Unauthenticated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := logger.Initialize("error")
			require.NoError(t, err, "Error init logger")

			s := &Server{
				Config: &config.ServerConfig{
					AuthenticationKey: secret,
				},
				Logger: l,
			}

			var ctx context.Context
			if tt.authToken != "" {
				md := metadata.New(map[string]string{AuthorizationMetadataKey: tt.authToken})
				ctx = metadata.NewIncomingContext(context.Background(), md)
			} else {
				ctx = context.Background()
			}

			req := &pb.AddTaskRequestV1{}
			info := &grpc.UnaryServerInfo{
				FullMethod: "/gophertask.gophertaskapi.v1.AccountServiceV1/AddAccount",
			}
			handler := func(ctx context.Context, req interface{}) (interface{}, error) { return nil, nil }

			if tt.exempted {
				info.FullMethod = "/gophertask.gophertaskapi.v1.AuthServiceV1/SignIn"
			}

			_, err = s.AuthInterceptor(ctx, req, info, handler)

			if tt.expectedStatus == codes.OK {
				assert.NoError(t, err, "expected no error")
			} else {
				assert.Error(t, err, "expected an error")
				assert.Equal(t, status.Code(err), tt.expectedStatus, "expected status code %v but got %v", tt.expectedStatus, status.Code(err))
			}
		})
	}
}
