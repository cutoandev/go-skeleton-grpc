package service

import (
  "x-user-service/pkg/database"
  "x-user-service/pkg/service/auth"
)

var mainService *MainService

type MainService struct {
  DBContext    database.DBContext
  TokenService auth.TokenService
}

func InitServices(dbContext database.DBContext) *MainService {
  mainService = registerServices(dbContext)
  return mainService
}

func registerServices(dbContext database.DBContext) *MainService {
  return &MainService{
    DBContext:    dbContext,
    TokenService: auth.TokenService{DBContext: dbContext},
  }
}
