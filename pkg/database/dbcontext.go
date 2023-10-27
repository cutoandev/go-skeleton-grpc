package database

import (
  "x-user-service/pkg/database/infrastructure"
  "x-user-service/pkg/database/models"
  "gorm.io/gorm"
)

type DBContext struct {
  DB       *gorm.DB
  User     infrastructure.IRepository[models.User]
  UserInfo infrastructure.IRepository[models.UserInfo]
}

func GetDBContext(dbConnection *gorm.DB) DBContext {
  return DBContext{
    DB:       dbConnection,
    User:     infrastructure.RepositoryBase[models.User]{DB: dbConnection},
    UserInfo: infrastructure.RepositoryBase[models.UserInfo]{DB: dbConnection},
  }
}
