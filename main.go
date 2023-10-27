package main

import (
  "fmt"
  "github.com/joho/godotenv"
  "google.golang.org/grpc"
  "log"
  "net"
  "x-user-service/bootstrap"
  "x-user-service/configs"
)

func main() {

  // Load environment variables
  err := godotenv.Load()
  if err != nil {
    log.Printf("Load env variables error occured. Err: %s", err)
  }

  // Get configs
  serverConfig := configs.GetServerConfigs()
  serverPort := serverConfig.Port
  if serverPort == "" {
    serverPort = "5001"
  }

  lis, errServer := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
  if errServer != nil {
    log.Fatalf("failed to listen: %v", errServer)
  }
  s := grpc.NewServer()

  // Register all controller
  bootstrap.RegisterServices(s)

  log.Printf("server listening at %v", lis.Addr())
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
