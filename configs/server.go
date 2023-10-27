package configs

import (
  "os"
  config "x-user-service/configs/models"
)

func GetServerConfigs() config.ServerConfig {
  return config.ServerConfig{
    Port: os.Getenv("SERVER_PORT"),
  }
}
