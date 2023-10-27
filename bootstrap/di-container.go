package bootstrap

import (
  "x-user-service/pkg/service"
  "x-user-service/pkg/database"
)

var diContainer *DIContainer

type DIContainer struct {
  DBContext database.DBContext
  Service   service.MainService
}

func NewDIContainer() *DIContainer {
  dbConnection := GetDBConnection()
  dbContext := database.GetDBContext(dbConnection)
  return &DIContainer{
    DBContext: database.GetDBContext(dbConnection),
    Service:   *service.InitServices(dbContext),
  }
}

func DIContainerInstance() *DIContainer {
  if diContainer == nil {
    return NewDIContainer()
  }
  return diContainer
}
