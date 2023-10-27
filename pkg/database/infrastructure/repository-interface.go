package infrastructure

import "gorm.io/gorm"

type IRepository[T any] interface {
  DBContext() *gorm.DB
  Model() *gorm.DB
  GetById(id interface{}) (T, error)
  GetAll() ([]T, error)
  Create(entity T, tx *gorm.DB) error
  Update(entity T, tx *gorm.DB, condition ...any) error
  Delete(entity T, tx *gorm.DB, condition ...any) error
  Where(condition ...any) ([]T, error)
}
