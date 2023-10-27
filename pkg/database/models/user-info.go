package models

import (
  "gorm.io/gorm"
)

type UserInfo struct {
  gorm.Model
  FirstName   string
  LastName    string
  MiddleNam   string
  DisplayName string
}
