package bootstrap

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
  "x-user-service/configs"
)

var dbContext *gorm.DB

func InitDatabaseConnection() {
  dbConfig := configs.GetDatabaseConfig()
  dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    dbConfig.DBHost,
    dbConfig.DBUser,
    dbConfig.DBPassword,
    dbConfig.DBName,
    dbConfig.DBPort,
  )

  // Open connection and check error
  var err error
  dbContext, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
  if err != nil {
    log.Panic("Error when open database connection")
  }
  log.Printf("database connect is successfully (host: %s)", dbConfig.DBHost)

}

func GetDBConnection() *gorm.DB {
  if dbContext != nil {
    return dbContext
  }
  InitDatabaseConnection()
  return dbContext
}
