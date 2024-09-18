package client

import (
	"context"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type AuthService interface {
	SignUp(login string, password string) (string, error)
	SignIn(login string, password string) (string, error)
}

type AuthServiceImpl struct {
	authClient pb.AuthServiceV1Client
}

func NewAuthService(authClient pb.AuthServiceV1Client) AuthService {
	return &AuthServiceImpl{
		authClient: authClient,
	}
}

func (s *AuthServiceImpl) SignUp(login string, password string) (string, error) {
	ctx := context.Background()
	req := &pb.SignUpRequestV1{Login: login, Password: password}
	resp, err := s.authClient.SignUp(ctx, req)
	if err != nil {
		return "", err
	}

	token := resp.Token
	return token, nil
}

func (s *AuthServiceImpl) SignIn(login string, password string) (string, error) {
	ctx := context.Background()
	req := pb.SignInRequestV1{Login: login, Password: password}
	resp, err := s.authClient.SignIn(ctx, &req)
	if err != nil {
		return "", err
	}

	token := resp.Token
	return token, nil
}
