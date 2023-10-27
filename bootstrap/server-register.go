package bootstrap

import (
  "google.golang.org/grpc"
  "x-user-service/protos/auth"
  authModel "x-user-service/controller/auth"
)

func RegisterServices(server *grpc.Server) {
  diContainer := DIContainerInstance()
  auth.RegisterAuthenticationServer(server, &authModel.AuthenticationImplementation{Service: diContainer.Service})
}
