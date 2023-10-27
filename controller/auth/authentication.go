package auth

import (
  "context"
  "x-user-service/protos/auth"
  "x-user-service/pkg/service"
)

type AuthenticationImplementation struct {
  auth.UnimplementedAuthenticationServer
  Service service.MainService
}

func (srv *AuthenticationImplementation) Token(ctx context.Context, in *auth.LoginPayload) (*auth.AuthenticationResponse, error) {
  // example code
  return srv.Service.TokenService.GenerateToken()
}
