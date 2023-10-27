package configs

import (
  "os"
  "x-user-service/configs/models"
)

// GetDatabaseConfig is a function will return all env related to database config
// Params:
//   - None
//
// Returns:
//   - DatabaseConfig: all config related to database config
func GetDatabaseConfig() config.DatabaseConfig {
  return config.DatabaseConfig{
    DBHost:     os.Getenv("DB_HOST"),
    DBName:     os.Getenv("DB_NAME"),
    DBUser:     os.Getenv("DB_USER"),
    DBPassword: os.Getenv("DB_PASSWORD"),
    DBPort:     os.Getenv("DB_PORT"),
  }
}
